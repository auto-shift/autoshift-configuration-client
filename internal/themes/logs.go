package interfaces

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/theme"

// )

// type logTheme struct{}

// var resourceImagePng = &fyne.StaticResource{
// 	StaticName: "image.png",
// 	StaticContent: []byte{}}

// var _ fyne.Theme = (*logTheme)(nil)

// func (lt logTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
// 	if name == theme.ColorNameBackground {
// 		if variant == theme.VariantLight {
// 			return color.White
// 		}
// 		return color.Black
// 	}

// 	return theme.DefaultTheme().Color(name, variant)
// }

// func (lt logTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
// 	if name == theme.IconNameHome {
// 		fyne.NewStaticResource("","")
// 	}

// 	return theme.DefaultTheme().Icon(name)
// }

// func (lt logTheme) Font(style fyne.TextStyle) fyne.Resource {
// 	return theme.DefaultTheme().Font(style)
// }

// func (lt logTheme) Size(name fyne.ThemeSizeName) float32 {
// 	return theme.DefaultTheme().Size(name)
// }
