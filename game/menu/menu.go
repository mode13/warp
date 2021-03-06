/*
Copyright (C) 2016 Andreas T Jonsson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package menu

import (
	"github.com/andreas-jonsson/nanovgo"
	"github.com/andreas-jonsson/warp/game"
)

type menuState struct {
}

func NewMenuState() *menuState {
	return &menuState{}
}

func (s *menuState) Name() string {
	return "menu"
}

func (s *menuState) Enter(from game.GameState, args ...interface{}) error {
	args[0].(game.GameControl).SwitchState("play", args[0])
	return nil
}

func (s *menuState) Exit(to game.GameState) error {
	return nil
}

func (s *menuState) Update(gctl game.GameControl) error {
	gctl.PollAll()
	return nil
}

func (s *menuState) Render(ctx *nanovgo.Context) error {
	return nil
}
