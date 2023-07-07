# IconChanger
IconChanger scans the folder for images and replaces them with the image you provide, keeping the file name and resolution.
It is useful for when you have a lot of images to replace, for example when you want to change the icon of an app.

## Usage

The first argument is the folder you want to scan, the second argument is the image you want to replace the images with.

```bash
go run src/main.go Assets.xcassets input.svg
```