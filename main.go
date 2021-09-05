package main

import (
	"fmt"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var frameCount int = 0
var frameSpeed int = 8 // how fast animations are?

func main() {
	screenHeight := 720
	screenWidth := 1280

	rl.InitWindow(screenWidth, screenHeight, "Trying things")
	rl.SetTargetFPS(60)

	camera := NewCustomCamera(180.0, 1.5, 10.0)

	rl.SetCameraMode(&camera.Camera, rl.CameraCustom)

	player := InitPlayer()
	// pixels, err := getSprites("sprites/blue.png", 24, 24)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	dt := rl.GetFrameTime()

	for !rl.WindowShouldClose() {
		dt = rl.GetFrameTime()
		frameCount++

		// Camera Movement
		if rl.IsKeyDown(rl.KeyDown) { // Zoom
			camera.Camera.FOVY += camera.ZoomSpeed * dt // TODO ZoomSpeed
			camera.Zoom += float64(camera.ZoomSpeed * dt)
		} else if rl.IsKeyDown(rl.KeyUp) {
			camera.Camera.FOVY -= camera.ZoomSpeed * dt
			camera.Zoom -= float64(camera.ZoomSpeed * dt)
		}

		if rl.IsKeyDown(rl.KeyPageUp) { // Up/Down
			camera.Angle.Y -= camera.RotationSpeed * dt
		} else if rl.IsKeyDown(rl.KeyPageDown) {
			camera.Angle.Y += camera.RotationSpeed * dt
		}

		if rl.IsKeyDown(rl.KeyRight) { // Rotate
			camera.Angle.X += camera.RotationSpeed * dt
		} else if rl.IsKeyDown(rl.KeyLeft) {
			camera.Angle.X -= camera.RotationSpeed * dt
		}

		player.PlayerMovement()
		camera.SetTarget(player.pos)

		camera.Update(dt)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera.Camera)

		rl.DrawGrid(20, 2.0)
		player.renderPlayer()
		//renderPixelsToCube(pixels, *player)

		rl.EndMode3D()

		//rl.DrawTexture(dinoTexture, screenWidth/2-dinoTexture.Width/2, screenHeight/2-dinoTexture.Height/2, rl.RayWhite)
		rl.DrawCube(rl.NewVector3(0.0, 0.1, 0.0), 32.0, 32.0, 32.0, rl.Red)
		//debug prints
		incr := 20
		start := 80 - incr
		incrY := func() int {
			start += 20
			return start
		}
		rl.DrawText(fmt.Sprintf("FPS: %v", rl.GetFPS()), 20, incrY(), 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Frame count: %v", frameCount), 20, incrY(), 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Camera X: %v", camera.Angle.X), 20, incrY(), 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Camera Y: %v", camera.Angle.Y), 20, incrY(), 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Camera POS: %v", camera.Camera.Position), 20, incrY(), 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Player state: %v", player.playerState), 20, incrY(), 20, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
