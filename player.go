package main

import (
	"log"

	rl "github.com/lachee/raylib-goplus/raylib"
)

type playerInfo struct {
	pos           rl.Vector3
	movementSpeed float32
	Size          rl.Vector3
	Texture       rl.Texture2D
	sprites       Sprite
	playerState   [2]string // current state, previous state
	currentFrame  int
}

func InitPlayer() *playerInfo {
	player := &playerInfo{}
	// default
	player.pos = rl.NewVector3(0.0, 0.0, 0.0)
	player.movementSpeed = 0.25
	size := float32(0.5)
	player.Size = rl.NewVector3(size, size, size)
	player.playerState[0] = "idle"

	var spriteState SpriteState
	spriteState.idle = []int{0, 1, 2, 3}
	spriteState.moving = []int{4, 5, 6, 7, 8, 9}
	sprites, err := getSprites("sprites/blue.png", 24, 24, spriteState)
	if err != nil {
		log.Println(err)
	}
	player.sprites = sprites
	player.currentFrame = 0

	return player
}

func (player *playerInfo) PlayerMovement() {
	// Left - Right
	isMoving := false
	if rl.IsKeyDown(rl.KeyA) {
		player.pos.X -= player.movementSpeed
		isMoving = true
	} else if rl.IsKeyDown(rl.KeyD) {
		player.pos.X += player.movementSpeed
		isMoving = true
	}

	// Up - Down
	if rl.IsKeyDown(rl.KeyW) {
		player.pos.Z -= player.movementSpeed
		isMoving = true
	} else if rl.IsKeyDown(rl.KeyS) {
		player.pos.Z += player.movementSpeed
		isMoving = true
	}

	if player.playerState[0] != player.playerState[1] {
		player.currentFrame = 0 // to show that we need to reset the animation cycle
	}
	player.playerState[1] = player.playerState[0]
	if isMoving {
		player.playerState[0] = "moving"
	} else {
		player.playerState[0] = "idle"
	}
}

func (player *playerInfo) renderPlayer() {
	log.Println(player.currentFrame)
	spriteAnimationIndex := []int{0}
	if player.playerState[1] == "idle" {
		spriteAnimationIndex = player.sprites.spriteState.idle
	} else if player.playerState[1] == "moving" {
		spriteAnimationIndex = player.sprites.spriteState.moving
	}
	log.Println(spriteAnimationIndex)

	if frameCount >= (60 / frameSpeed) {
		frameCount = 0

		player.currentFrame++
		if player.currentFrame > len(spriteAnimationIndex)-1 {
			player.currentFrame = 0
		}
	}
	getSprite := player.sprites.sprites[spriteAnimationIndex[player.currentFrame]]
	renderPixelsToCube(getSprite, *player)
}

func renderPixelsToCube(pixels [][]Pixel, obj playerInfo) {
	//startY := (obj.pos.Y - ((float32(len(pixels)) * obj.Size.Y) / 2))    // start pos to render Y
	startY := float32(0.0)
	startX := (obj.pos.X - ((float32(len(pixels[0])) * obj.Size.X) / 2)) // start pos to render X

	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[y]); x++ {
			if pixels[y][x].A > 0 { // if alpha is 0, why bother
				pixelColour := rl.NewColor(uint8(pixels[y][x].R), uint8(pixels[y][x].G), uint8(pixels[y][x].B), uint8(pixels[y][x].A))
				posX := float32(startX + (obj.Size.X * float32(x)))
				posY := float32(startY + (obj.Size.Y * float32(y)))
				rl.DrawCube(rl.NewVector3(posX, posY, obj.pos.Z), obj.Size.X, obj.Size.Y, obj.Size.Z, pixelColour)
				//rl.DrawCubeWires(rl.NewVector3(posX, posY, obj.pos.Z), obj.Size.X, obj.Size.Y, obj.Size.Z, rl.Black)
			}
		}
	}

}
