package main

import (
	rl "github.com/lachee/raylib-goplus/raylib"
)

func main() {
	screenWidth := 1280
	screenHeight := 720

	rl.InitWindow(screenWidth, screenHeight, "Trying things")
	rl.SetTargetFPS(60)

	camera := NewCustomCamera(180.0, 3.0, 100.0)
	rl.SetCameraMode(&camera.Camera, rl.CameraCustom) // Set a first person misc.CustomCamera mode

	dino := rl.LoadImage("sprites/dino.png")
	rl.ImageResizeNN(dino, int(dino.Width*16), int(dino.Height*16))
	dinoTexture := rl.LoadTextureFromImage(dino)

	dt := rl.GetFrameTime()
	camera.SetTarget(rl.NewVector3(0.0, 0.0, 0.0))
	camera.Update(dt)

	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()

		if rl.IsKeyDown(rl.KeyRight) { // Rotate
			camera.Angle.X += camera.RotationSpeed * dt
		} else if rl.IsKeyDown(rl.KeyLeft) {
			camera.Angle.X -= camera.RotationSpeed * dt
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera.Camera)

		rl.DrawCubeTexture(dinoTexture, rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.White)

		rl.DrawGrid(20, 2.0)
		rl.EndMode3D()

		//rl.DrawTexture(dinoTexture, screenWidth/2-dinoTexture.Width/2, screenHeight/2-dinoTexture.Height/2, rl.RayWhite)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
