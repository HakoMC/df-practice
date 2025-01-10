package dfpractice

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

type MoveHandler struct {
	player.NopHandler // Embed the NopHandler for default behavior
	world             *world.World
}

func NewMoveHandler(w *world.World) *MoveHandler {
	return &MoveHandler{world: w}
}

func (h MoveHandler) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	p := ctx.Val()
	_ = p
	fmt.Println("Pos", newPos, "Rot", newRot)
	h.world.Exec(func(tx *world.Tx) {
		// Use tx to edit the world, for example:
		tx.PlaySound(newPos, sound.Experience{})
		blockPos := cube.PosFromVec3(newPos.Add(mgl64.Vec3{0, -1, 0})) // mgl64.Vec3 を cube.Pos に変換
		tx.SetBlock(blockPos, block.Dirt{}, nil)                       // 正しい引数を渡す。nilはデフォルトオプション
	})
}
