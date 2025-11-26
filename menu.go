package main

import (
	"fmt"
	"image/color"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/shnm"

	"github.com/SWITCHSCIENCE/ffb_steering_controller/settings"
	"github.com/SWITCHSCIENCE/picossci-ffb-wheel/board"
)

const maxVisibleItems = 3

type Value struct {
	Min, Max, Step float32
	Value          float32
	Format         string
}

func (v *Value) String() string {
	return fmt.Sprintf(v.Format, v.Value)
}

type Item struct {
	Name  string
	Value *Value
}

type Menu struct {
	settingMode bool
	edit        bool
	index       int
	offset      int
	params      []*Item
	current     settings.Settings
}

func NewMenu() *Menu {
	return &Menu{
		params: []*Item{
			{"Exit without Saving", nil},
			{"Save and Exit", nil},
			{"Lock2Lock", &Value{Min: 180, Max: 1440, Step: 90, Format: "%.0f"}},
			{"NeutralAdjust", &Value{Min: -10, Max: 10, Step: 0.5, Format: "%.1f"}},
			{"MaxCenteringForce", &Value{Min: 0, Max: 1000, Step: 100, Format: "%.0f"}},
			{"CoggingTorqueCancel", &Value{Min: 0, Max: 256, Step: 1, Format: "%.0f"}},
			{"Viscosity", &Value{Min: 0, Max: 1024, Step: 16, Format: "%.0f"}},
			{"SoftLockForce", &Value{Min: 0, Max: 16, Step: 1, Format: "%.0f"}},
		},
	}
}

func (m *Menu) IsSetting() bool {
	return m.settingMode
}

func (m *Menu) Up() {
	if !m.settingMode {
		return
	}
	if m.edit {
		item := m.params[m.index]
		if item.Value == nil {
			return
		}
		item.Value.Value -= item.Value.Step
		if item.Value.Value < item.Value.Min {
			item.Value.Value = item.Value.Min
		}
		m.ShowEdit(item)
	} else {
		if m.index > 0 {
			m.index--
			if m.index < m.offset {
				m.offset = m.index
			}
			m.ShowMenu()
		}
		return
	}
}

func (m *Menu) Down() {
	if !m.settingMode {
		return
	}
	if m.edit {
		item := m.params[m.index]
		if item.Value == nil {
			return
		}
		item.Value.Value += item.Value.Step
		if item.Value.Value > item.Value.Max {
			item.Value.Value = item.Value.Max
		}
		m.ShowEdit(item)
	} else {
		if m.index < len(m.params)-1 {
			m.index++
			if m.index-m.offset >= maxVisibleItems {
				m.offset++
			}
			m.ShowMenu()
		}
		return
	}
}

func (m *Menu) Enter() {
	if !m.settingMode {
		m.settingMode = true
		m.index = 0
		m.offset = 0
		m.current = settings.Get()
		m.params[2].Value.Value = float32(m.current.Lock2Lock)
		m.params[3].Value.Value = m.current.NeutralAdjust
		m.params[4].Value.Value = float32(m.current.MaxCenteringForce)
		m.params[5].Value.Value = float32(m.current.CoggingTorqueCancel)
		m.params[6].Value.Value = float32(m.current.Viscosity)
		m.params[7].Value.Value = float32(m.current.SoftLockForceMagnitude)
		m.ShowMenu()
		return
	}
	switch m.params[m.index].Name {
	case "Save and Exit":
		// TODO: persistent params
		b := settings.Marshal(m.current)
		if err := writeFlashBlock(b); err != nil {
			println(err)
		}
		fallthrough
	case "Exit without Saving":
		m.settingMode = false
		board.LCD.Show(board.Logo)
		board.LCD.Display()
		return
	default:
		if m.edit {
			m.edit = false
			m.ShowMenu()
		} else {
			m.edit = true
			m.ShowEdit(m.params[m.index])
		}
	}
}

func (m *Menu) ShowEdit(p *Item) {
	switch p.Name {
	case "Lock2Lock":
		m.current.Lock2Lock = int32(p.Value.Value)
	case "NeutralAdjust":
		m.current.NeutralAdjust = p.Value.Value
	case "MaxCenteringForce":
		m.current.MaxCenteringForce = int32(p.Value.Value)
	case "CoggingTorqueCancel":
		m.current.CoggingTorqueCancel = int32(p.Value.Value)
	case "Viscosity":
		m.current.Viscosity = int32(p.Value.Value)
	case "SoftLockForce":
		m.current.SoftLockForceMagnitude = int32(p.Value.Value)
	}
	settings.Update(m.current)
	lines := []string{fmt.Sprintf("%s:", p.Name), "", fmt.Sprintf("    %s", p.Value.String()), ""}
	board.LCD.Clear()
	for i, line := range lines {
		tinyfont.WriteLine(board.LCD, &shnm.Shnmk12, 1, int16(12*(i+1)), line, color.RGBA{0, 0, 0, 255})
	}
	board.LCD.Display()
}

func (m *Menu) ShowMenu() {
	lines := []string{fmt.Sprintf("=== Setting(%d/%d) ===", m.index+1, len(m.params))}
	for idx, param := range m.params[m.offset : m.offset+maxVisibleItems] {
		c := ' '
		if m.index == idx+m.offset {
			c = '>'
		}
		lines = append(lines, fmt.Sprintf("%c%s", c, param.Name))
	}
	board.LCD.Clear()
	for i, line := range lines {
		tinyfont.WriteLine(board.LCD, &shnm.Shnmk12, 1, int16(12*(i+1)), line, color.RGBA{0, 0, 0, 255})
	}
	board.LCD.Display()
}
