// SPDX-License-Identifier: MIT

// Package slogadapter bridges *slog.Logger callers (notably
// github.com/prometheus/exporter-toolkit) onto the package-level logrus
// logger used by the rest of junos_exporter. Records keep their original
// level and structured attributes instead of being collapsed into a
// pre-rendered text line.
package slogadapter

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	log "github.com/sirupsen/logrus"
)

// New returns a *slog.Logger backed by a Handler that forwards to the
// package-level logrus logger.
func New() *slog.Logger {
	return slog.New(&Handler{})
}

// Handler implements slog.Handler. The zero value is ready to use.
type Handler struct {
	attrs  []slog.Attr
	groups []string
}

// Enabled lets logrus's current level filter slog records before any work
// (attr collection, string formatting) is done.
func (h *Handler) Enabled(_ context.Context, lvl slog.Level) bool {
	return mapLevel(lvl) <= log.GetLevel()
}

// Handle forwards r to logrus at the corresponding level, with all attrs
// (those set via WithAttrs plus those on the record) as logrus fields.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	fields := log.Fields{}
	for _, a := range h.attrs {
		addAttr(fields, h.groups, a)
	}
	r.Attrs(func(a slog.Attr) bool {
		addAttr(fields, h.groups, a)
		return true
	})

	entry := log.WithFields(fields)
	switch mapLevel(r.Level) {
	case log.ErrorLevel:
		entry.Error(r.Message)
	case log.WarnLevel:
		entry.Warn(r.Message)
	case log.DebugLevel:
		entry.Debug(r.Message)
	default:
		entry.Info(r.Message)
	}
	return nil
}

// WithAttrs returns a new Handler whose Handle calls will include attrs in
// every emitted record.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	clone := h.clone()
	clone.attrs = append(clone.attrs, attrs...)
	return clone
}

// WithGroup returns a new Handler that prefixes the keys of subsequent
// attrs with name (joined by dots, per slog conventions).
func (h *Handler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	clone := h.clone()
	clone.groups = append(clone.groups, name)
	return clone
}

func (h *Handler) clone() *Handler {
	return &Handler{
		attrs:  append([]slog.Attr(nil), h.attrs...),
		groups: append([]string(nil), h.groups...),
	}
}

func mapLevel(lvl slog.Level) log.Level {
	switch {
	case lvl >= slog.LevelError:
		return log.ErrorLevel
	case lvl >= slog.LevelWarn:
		return log.WarnLevel
	case lvl >= slog.LevelInfo:
		return log.InfoLevel
	default:
		return log.DebugLevel
	}
}

func addAttr(fields log.Fields, groups []string, a slog.Attr) {
	v := a.Value.Resolve()
	if a.Key == "" && v.Kind() != slog.KindGroup {
		return
	}

	if v.Kind() == slog.KindGroup {
		groupAttrs := v.Group()
		if len(groupAttrs) == 0 {
			return
		}
		next := groups
		if a.Key != "" {
			next = append(append([]string(nil), groups...), a.Key)
		}
		for _, child := range groupAttrs {
			addAttr(fields, next, child)
		}
		return
	}

	key := a.Key
	if len(groups) > 0 {
		key = strings.Join(groups, ".") + "." + key
	}
	fields[key] = valueAny(v)
}

func valueAny(v slog.Value) any {
	if v.Kind() == slog.KindLogValuer {
		return valueAny(v.Resolve())
	}
	switch v.Kind() {
	case slog.KindString:
		return v.String()
	case slog.KindInt64:
		return v.Int64()
	case slog.KindUint64:
		return v.Uint64()
	case slog.KindFloat64:
		return v.Float64()
	case slog.KindBool:
		return v.Bool()
	case slog.KindDuration:
		return v.Duration()
	case slog.KindTime:
		return v.Time()
	case slog.KindAny:
		return v.Any()
	}
	return fmt.Sprint(v.Any())
}
