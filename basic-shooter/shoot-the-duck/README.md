# Final Version

Final version of the "Shoot the duck" game. The main difference from the
[exercise3](https://github.com/DiabeticOwl/Learning_Ebiten/blob/master/basic-shooter/exercise3/)
is a code refactor in which the project is divided in three packages; `main`,
`object` and `utils`, alongside the already implemented `resources` package.

## Main

Is the heart of the game in which the [framework's engine](https://ebitengine.org/)
is instantiated and run looping through a list of objects that will draw and
update themselves accordingly.

## Object

Describes the structure of each object that will appear and interact with the
user in the game.

## Util

Corpora of functions that automatizes processes like the decoding of an image in
a slice of bytes.

## Resources

Collection of miscellaneous files used throughout the entire application.

![gamePreview](https://github.com/DiabeticOwl/Learning_Ebiten/blob/master/basic-shooter/shoot-the-duck/Shoot-The-Duck-Initial.gif)
