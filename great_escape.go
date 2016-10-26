// Great escape


package main

import (
	"container/list"
	"fmt"
	"os"
)

type wall struct {
	x, y int
	or   string
}
type cord struct {
	x, y int
}

type player struct {
	pt     cord
	nWalls int
}

type PathList []cord

type traverseFunc func([]cord) []cord

var w, h, myId int
var x, y int
var walls []wall
var players []player
var paths []PathList
var lastMove cord
var fnRegistry = map[string]traverseFunc{
	"DOWN":  func(ptList []cord) []cord { return canGoDown(ptList) },
	"UP":    func(ptList []cord) []cord { return canGoUp(ptList) },
	"LEFT":  func(ptList []cord) []cord { return canGoLeft(ptList) },
	"RIGHT": func(ptList []cord) []cord { return canGoRight(ptList) },
}

func main() {
	// w: width of the board
	// h: height of the board
	// playerCount: number of players (2 or 3)
	// myId: id of my player (0 = 1st player, 1 = 2nd player, ...)
	var playerCount int
	fmt.Scan(&w, &h, &playerCount, &myId)
	// 	fmt.Fprintln(os.Stderr, "h: ", h, "w:", w)

	// var odir string
	// var oX,oY int
	walls = make([]wall, 0)
	players = make([]player, 0)

	// 	lastCord.x = -1
	// 	lastCord.y = -1

	for {
		for i := 0; i < playerCount; i++ {
			if i == 0 {
				players = make([]player, 0)
			}
			// x: x-coordinate of the player
			// y: y-coordinate of the player
			// wallsLeft: number of walls available for the player
			var x, y, wallsLeft int
			fmt.Scan(&x, &y, &wallsLeft)
			// 			fmt.Fprintln(os.Stderr, "i: ", i, "x:", x, " y:", y, " WL:", wallsLeft)
			p := player{pt: cord{x: x, y: y}, nWalls: wallsLeft}
			players = append(players, p)
		}
		// wallCount: number of walls on the board
		var wallCount int
		fmt.Scan(&wallCount)
		allPaths := make([]PathList, 0, 0)

		for i := 0; i < wallCount; i++ {
			var wallX, wallY int
			var wallOrientation string
			fmt.Scan(&wallX, &wallY, &wallOrientation)
			fmt.Fprintln(os.Stderr, wallX, wallY, wallOrientation)
			wall := wall{x: wallX, y: wallY, or: wallOrientation}
			walls = append(walls, wall)
		}
		for id, player := range players {
			p := (process(player.pt, id))
			allPaths = append(allPaths, p)
		}

		fmt.Println(wallOrMove(allPaths))
	}
}

func wallOrMove(paths []PathList) string {
	var m string
	for id, _ := range players {
		fmt.Fprintln(os.Stderr,id, paths[id])
	}
	mysteps := len(paths[myId])
	if players[myId].nWalls > 0 {
		if len(paths) == 2 {
			// 			if !(len(paths[0]) > 5 && len(paths[1]) > 5) {
// 			fmt.Fprintln(os.Stderr, mysteps)

			if myId==0{
			    if  len(paths[1]) < 6 {
				    m = block(paths[1], 1)
			    }
			}else {
			    if  len(paths[0]) < 6 {
			    	m = block(paths[0], 0)
		    	}
			}


			 //else
// 			else if mysteps <= 6 {
// 				//go defensive
// 				m = saveMyAss(paths[myId])
// 			}
			// 			}
		} else if len(paths) == 3 {
			zero := len(paths[0])
			one := len(paths[1])
			two := len(paths[2])

			if players[0].pt.x == -1 {
				zero = 1000
			}
			if players[1].pt.x == -1 {
				one = 1000
			}
			if players[2].pt.x == -1 {
				two = 10000
			}
			if !(one > 5 && two > 5 && zero > 5) {
				switch myId {
				case 0:
					if mysteps > one {
						if one <= two {
							m = block(paths[1], 1) //1)
						} else {
							m = block(paths[2], 2)
						}
					}
				case 1:
					if mysteps > zero {
						if zero <= two {
							m = block(paths[0], 0)
						} else {
							m = block(paths[2], 2) //2)
						}
					}

				case 2:
					if mysteps > one {
						if zero <= one {
							m = block(paths[0], 0) //0)
						} else {
							m = block(paths[1], 1)
						}
					}
				default:

				}
			}
		}
	}
	if len(m) > 0 {
		return m
	}
	return move(paths[myId][0], paths[myId][1])
}

func saveMyAss(paths []cord) string {
// 	fmt.Fprintln(os.Stderr, paths)
	m := isbackCovered(paths)
	// 	if !isbackCovered(lastMove, paths[0]) {
	// 		m := move(lastMove, paths[0])
	// 		fmt.Fprintln(os.Stderr, "saveMyAss: ", m)

	// 		var w string
	// 		if m == "UP" {
	// 			w = canPlaceWall(paths[0], "H")
	// 		} else if m == "DOWN" {
	// 		} else if m == "LEFT" {
	// 		} else if m == "RIGHT" {
	// 			//fmt.Fprintln(os.Stderr, "canPlaceWall: ", paths[0], "V")
	// 			w = canPlaceWall(paths[0], "V")
	// 		}
	// 		if len(w) > 0 {
	// 			fmt.Fprintln(os.Stderr, "Wall", w)
	// 			return w
	// 		}
	if len(m) > 0 {
		return m
	} else {
		w := coverTopBottom(paths)
		if len(w) > 0 {
			fmt.Fprintln(os.Stderr, "Wall", w)
			return w
		}
	}
	return ""
}
func coverTopBottom(paths []cord) string {
	fmt.Fprintln(os.Stderr, "coverTopBottom: ", paths)
	var w string
	w = ""
	for k, _ := range paths {
		if k == 0 {
			continue
		}

		m := move(paths[k-1], paths[k])
		fmt.Fprintln(os.Stderr, "coverTopBottom: ", k, m)

		if m == "UP" {
			w = canPlaceWall(paths[k], "H")
		} else if m == "DOWN" {
		} else if m == "LEFT" {
		} else if m == "RIGHT" {
			fmt.Fprintln(os.Stderr, "coverTopBottom: canPlaceWall: ", paths[k], "H")
			w = canPlaceWall(paths[k-1], "H")
			fmt.Fprintln(os.Stderr, "coverTopBottom: ", k, m, w)
			if len(w) > 0 {
				return w
			}

		}
	}
	return ""
}
func isbackCovered(paths []cord) string {
	fmt.Fprintln(os.Stderr, "isbackCovered")
	// 	return false

	m := move(lastMove, paths[0])
	fmt.Fprintln(os.Stderr, "saveMyAss: ", m)

	var w string
	if m == "UP" {
		w = canPlaceWall(paths[0], "H")
	} else if m == "DOWN" {
	} else if m == "LEFT" {
	} else if m == "RIGHT" {
		fmt.Fprintln(os.Stderr, "canPlaceWall: ", paths[0], "V")
		w = canPlaceWall(paths[0], "V")
		if len(w) > 0 {
			return w
		} else {
			c := cord{x: paths[0].x, y: paths[0].y - 1}
			w = canPlaceWall(c, "V")
		}
	}
	return w
}

func block(paths []cord, index int) string {
	n := 0
// 	switch index {
// 	case 0:
// 		//right
// 		for k, _ := range paths {
// 			if k > 0 {
// 				if paths[k-1].x != paths[k].x {
// 					n = k - 1
// 					break
// 				}
// 			}
// 		}
// 	case 1:
// 		for k, _ := range paths {
// 			if k > 0 {
// 				if paths[k-1].x != paths[k].x {
// 					n = k - 1
// 					break
// 				}
// 			}
// 		}
// 	case 2:
// 		for k, _ := range paths {
// 			if k > 0 {
// 				if paths[k-1].y != paths[k].y {
// 					n = k - 1
// 					break
// 				}
// 			}
// 		}
// 	}
	fmt.Fprintln(os.Stderr, "in block...", paths, n)
	// 		first := paths[len(paths)-1]
	// 		second := paths[len(paths)-2]
	first := paths[n]
	second := paths[n+1]
	var ctr int
	if first.x == second.x {
		if first.y < second.y {
			ctr = 1
		} else {
			ctr = 0
		}
		//horizontal wall
		pt1 := cord{x: first.x, y: first.y + ctr}
		wall := canPlaceWall(pt1, "H")
		if len(wall) <= 0 {
			pt2 := cord{x: first.x - 1, y: first.y + ctr}
			wall = canPlaceWall(pt2, "H")
		}
		if len(wall) > 0 {
			return wall
		}
	} else if first.y == second.y {
		if first.x < second.x {
			ctr = 1
		} else {
			ctr = 0
		}
		//vertical wall
		pt1 := cord{x: first.x + ctr, y: first.y}
		wall := canPlaceWall(pt1, "V")
		if len(wall) <= 0 {
			pt2 := cord{x: first.x + ctr, y: first.y - 1}
			wall = canPlaceWall(pt2, "V")
		}
		if len(wall) > 0 {
			return wall
		}
	}
	// 	tmpPath:=paths[:len(paths)-1]
// 	tmpPath := paths[1:]
// 	fmt.Fprintln(os.Stderr, "rec in block...", tmpPath, len(tmpPath))

// 	if len(tmpPath) >= 2 {
// 		return block(tmpPath, len(tmpPath)-2)
// 	} else {
		return ""
// 	}
}
func canPlaceWall(c cord, ori string) string {
	if (ori == "H" && (c.x > w-2 || c.x < 0)) || (ori == "V" && (c.y > h-2 || c.y < 0)) {
		// 		fmt.Fprintln(os.Stderr, "no wall", ori, c)
		return ""
	}
	for _, w := range walls {
		// 		if w.or == ori {
		// 	fmt.Fprintln(os.Stderr, "walls", w,"pt:",c)
		if ori == "H" && w.or == "H" {
			if (w.y == c.y) && (w.x == c.x || w.x == c.x-1 || w.x == c.x+1) {
				// 	fmt.Fprintln(os.Stderr, "no wall", ori, c)
				return ""
			}
		} else if ori == "H" && w.or == "V" {
			// fmt.Fprintln(os.Stderr, "wall?", ori, c,w)
			if c.x == w.x-1 && c.y-1 == w.y {
				// 	fmt.Fprintln(os.Stderr, "no wall", ori, c)
				return ""
			}
		} else if ori == "V" && w.or == "H" {
// 			fmt.Fprintln(os.Stderr, "wall?", ori, c, w)
			if w.x == c.x-1 && w.y == c.y+1 {
				// fmt.Fprintln(os.Stderr, "no wall", ori, c)
				return ""
			}
		} else if ori == "V" && w.or == "V" {
			if (w.x == c.x) && (w.y == c.y || w.y == c.y-1 || w.y == c.y+1) {
				// 	fmt.Fprintln(os.Stderr, "no wall", ori, c)
				return ""
			}
		}
		// 		}
	}
	wall := wall{x: c.x, y: c.y, or: ori}
	walls = append(walls, wall)
	for id, player := range players {
		p := (process(player.pt, id))
		if p == nil {
			return ""
		}
	}
	walls = walls[:len(walls)-1]


	return fmt.Sprintf("%d %d %s", c.x, c.y, ori)
}

func process(pt cord, id int) []cord {
	var pathList list.List
	path := make([]cord, 0, 0)
	path = append(path, cord{x: pt.x, y: pt.y})
	pathList.PushBack(path)

	for pathList.Len() > 0 {
		ptList := pathList.Front().Value.([]cord)
		if ptList == nil || len(ptList) == 0 {
			break
		}
		if reached((ptList[len(ptList)-1]), id) {
			// 			fmt.Fprintln(os.Stderr, "REached", id, ptList)
			return ptList
		}

		for _, v := range fnRegistry {
			child := v(ptList)
			if child != nil && len(child) != 0 {
				pathList.PushBack(child)
			}
		}
		ptList = pathList.Front().Value.([]cord)
		pathList.Remove(pathList.Front())
	}
	fmt.Fprintln(os.Stderr, "NOT Possible")

	return nil

}

func reached(pt cord, id int) bool {
	switch id {
	case 0:
		if pt.x == w-1 {
			return true
		}
	case 1:
		if pt.x == 0 {
			return true
		}
	case 2:
		if pt.y == h-1 {
			return true
		}
	default:
		return false
	}
	return false
}

func move(first cord, second cord) string {
	// fmt.Println("what???")
	lastMove.x = first.x
	lastMove.y = first.y

	if first.x == second.x {
		if first.y < second.y {
			return "DOWN"
		} else {
			return "UP"
		}
	} else if first.y == second.y {
		if first.x < second.x {
			return "RIGHT"
		} else {
			return "LEFT"
		}
	}
	return "LEFT"
}

func canGoRight(ptList []cord) []cord {
	// 	fmt.(os.Stderr,"walls", walls)
	if ptList == nil || len(ptList) == 0 {
		return nil
	}
	c := ptList[len(ptList)-1]
	if c.x == w-1 {
		return nil
	}
	for _, w := range walls {
		if w.or == "V" {
			// 	fmt.(os.Stderr, "walls", w,"pt:",c)
			if (w.x == c.x+1) && (w.y == c.y || w.y == c.y-1) {
				// fmt.(os.Stderr, "walls", w, "pt:", c)
				return nil
			}
		}
	}

	newPath := copyList(ptList)
	child := cord{x: c.x + 1, y: c.y}
	for _, v := range ptList {
		if v.x == child.x && v.y == child.y {
			return nil
		}
	}
	newPath = append(newPath, child)
	return newPath
}

func canGoLeft(ptList []cord) []cord {
	if ptList == nil || len(ptList) == 0 {
		return nil
	}
	c := ptList[len(ptList)-1]
	if c.x == 0 {
		//on extring letft
		return nil
	}
	for _, w := range walls {
		// fmt.(os.Stderr, "walls", w)

		if w.or == "V" {
			if (w.x == c.x) && (w.y == c.y || w.y == c.y-1) {
				return nil
			}
		}
	}
	// fmt.(os.Stderr, "Can go Left...!!!!")
	newPath := copyList(ptList)
	child := cord{x: c.x - 1, y: c.y}
	for _, v := range ptList {
		if v.x == child.x && v.y == child.y {
			return nil
		}
	}
	newPath = append(newPath, child)
	return newPath
}

func canGoUp(ptList []cord) []cord {
	//      fmt.(os.Stderr,"walls", walls)
	if ptList == nil || len(ptList) == 0 {
		return nil
	}
	c := ptList[len(ptList)-1]
	if c.y == 0 {
		return nil
	}
	for _, w := range walls {
		//fmt.(os.Stderr, "walls", w)
		if w.or == "H" {
			if (w.y == c.y) && (w.x == c.x || w.x == c.x-1) {
				// fmt.(os.Stderr, "WALL! Cant go up...!!!!")
				return nil
			}
		}
	}
	// fmt.(os.Stderr, "Can go Up...!!!!")
	newPath := copyList(ptList)
	child := cord{x: c.x, y: c.y - 1}
	for _, v := range ptList {
		if v.x == child.x && v.y == child.y {
			return nil
		}
	}
	newPath = append(newPath, child)
	return newPath
}

func canGoDown(ptList []cord) []cord {
	//      fmt.(os.Stderr,"walls", walls)
	if ptList == nil || len(ptList) == 0 {
		return nil
	}
	c := ptList[len(ptList)-1]
	if c.y == h-1 {
		// fmt.(os.Stderr, "Reched down")
		return nil
	}
	// fmt.Println("current pt: ", c)
	for _, w := range walls {
		//fmt.(os.Stderr, "walls", w)
		if w.or == "H" {
			if (w.y == c.y+1) && (w.x == c.x || w.x == c.x-1) {
				// fmt.(os.Stderr, "WALL! Cant go Down...!!!!")
				return nil
			}
		}
	}
	// fmt.(os.Stderr, "go Down...!!!!")
	newPath := copyList(ptList)
	child := cord{x: c.x, y: c.y + 1}
	for _, v := range ptList {
		if v.x == child.x && v.y == child.y {
			return nil
		}
	}
	newPath = append(newPath, child)
	return newPath
}

func copyList(list []cord) []cord {
	newPath := make([]cord, 0, 0)
	for _, v := range list {
		// fmt.Println("value copied: ", v)
		newPath = append(newPath, v)
	}
	return newPath
}
