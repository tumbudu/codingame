//Back to the Code

package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
type Data[20][35]int

type cord struct {
    x, y int
}

type Rect struct{
    ltpt, rbpt cord
    area int
}

func main() {
    // opponentCount: Opponent count
    var opponentCount int
    // var downDir bool

    // var up, right, moveCtr int
    // up = -1
    // right = -1
    // moveCtr=0

    // var newX, newY int
    // var last Rect
    // var oVisited bool
    var mypath[] cord
    fmt.Scan(&opponentCount)
    mypath = make([]cord,0,0)
    // var goingStart bool
    for {
        var gameRound int
        fmt.Scan(&gameRound)
        fmt.Fprintln(os.Stderr, gameRound)

        // x: Your x position
        // y: Your y position
        // backInTimeLeft: Remaining back in time
        var x, y, backInTimeLeft int
        fmt.Scan(&x, &y, &backInTimeLeft)
	    fmt.Fprintln(os.Stderr,"x,y,backintime", x, y, backInTimeLeft)

        for i := 0; i < opponentCount; i++ {
            // opponentX: X position of the opponent=
            // opponentY: Y position of the opponent
            // opponentBackInTimeLeft: Remaining back in time of the opponent
            var opponentX, opponentY, opponentBackInTimeLeft int
            fmt.Scan(&opponentX, &opponentY, &opponentBackInTimeLeft)
            fmt.Fprintln(os.Stderr, opponentX, opponentY, opponentBackInTimeLeft)

        }

        // mapState := scanMap(mypath)
        // // print(mapState)
        // rectList := getRectList(mapState)
        // rect := getTop(rectList)
        // fmt.Fprintln(os.Stderr, rect)
        mypath = isBackinTime(mypath,cord{x,y})
        mapState := scanMap(mypath)
        // print(mapState)
        rectList := getRectList(mapState)
        rect := getTop(rectList, cord{x,y})
        fmt.Fprintln(os.Stderr, rect)
        // if last.area == 0{
        //     last = rect
        //     goingStart = true
        // }
        pt := getNextPt(rect, cord{x,y}, mypath)
        fmt.Println(pt.x, " ", pt.y)
        mypath = append(mypath, pt)

        // if last.area == 0{
        //     last = rect
        //     goingStart = true
        // }


        // if goingStart {
        //     //rect origin changed.. change last..
        //     if !(rect.ltpt.x == last.ltpt.x && rect.ltpt.y == last.ltpt.y){
        //         fmt.Fprintln(os.Stderr,"rect start changed...", rect, last)
        //         last = rect
        //     }
        //     fmt.Println(last.ltpt.x," ", last.ltpt.y)
        //     // mypath = append(mypath, cord{x,y})
        //     fmt.Fprintln(os.Stderr,"going left top", mypath)

        //     if(x==last.ltpt.x && y==last.ltpt.y){
        //         fmt.Fprintln(os.Stderr,"visited left top.. now goto right bottom...")
        //         goingStart = false
        //         if oVisited == true{
        //             fmt.Fprintln(os.Stderr,"Rect complete",mypath)
        //             last.area=0
        //             mypath=mypath[:0]
        //             oVisited = false
        //         }else{
        //             oVisited = true
        //         }
        //     }
        // }else{
        //     if rect.rbpt.x != last.ltpt.x && rect.rbpt.y != last.ltpt.y{
        //         last = rect
        //     }
        //     fmt.Println(last.rbpt.x," ", last.rbpt.y)
        //     // mypath = append(mypath, cord{x,y})
        //     fmt.Fprintln(os.Stderr,"going right bottam", mypath)

        //     if(x==last.rbpt.x && y==last.rbpt.y){
        //         fmt.Fprintln(os.Stderr,"visited left top.. now goto right bottom...")
        //         oVisited = true
        //         goingStart = true
        //     }
        // }
    }
}


func isBackinTime(mypath []cord, pt cord) []cord{
    if len(mypath) > 0{
        lastPos:=mypath[len(mypath)-1]
        fmt.Fprintln(os.Stderr,pt,lastPos)
        if lastPos.x == pt.x && lastPos.y ==pt.y{
            fmt.Fprintln(os.Stderr,"No back movement")
        }else{
            fmt.Fprintln(os.Stderr,"back movement happended....")
            var index int
            for k,v:=range(mypath){
                fmt.Fprintln(os.Stderr,k,v,pt)
                if v.x==pt.x && v.y == pt.y{
                    index = k
                    break
                }
            }
            fmt.Fprintln(os.Stderr,">>>>>>>>>>>>>",index,mypath[:index])
            return mypath[:index]
        }
    }
    return mypath
}

func getNextPt(rect Rect, currentPos cord, mypath[] cord) cord{
    path := getPath(rect)
    pt := move(path, currentPos, mypath)
    return pt
}

func getPath(rect Rect)[]cord{
    path := make([]cord,0,0)
    x1:= rect.ltpt.x
    y1:= rect.ltpt.y

    x2:= rect.rbpt.x
    y2:= rect.rbpt.y

    for i := x1; i <= x2; i++ {
        pt:=cord{i,y1}
        path = append(path, pt)
    }
    for i := y1; i <= y2; i++ {
        pt:=cord{x2,i}
        path = append(path, pt)
    }
    for i := x2; i >= x1; i-- {
        pt:=cord{i,y2}
        path = append(path, pt)
    }
    for i := y2; i >= y1; i-- {
        pt:=cord{x1,i}
        path = append(path, pt)
    }

    return path
}

func move(path[] cord, pt cord, mypath[] cord) cord{
    fmt.Fprintln(os.Stderr,"path: >>>>", path)
    fmt.Fprintln(os.Stderr,"mapath: >>>>", mypath)
    minDist := 100
    var nearestPt cord
    newpaht := make([]cord,0,0)
    var found,onpath bool
    for _,v:=range path{
        // fmt.Fprintln(os.Stderr,"New path:", newpaht)
        if v.x==pt.x&& v.y==pt.y {
            onpath = true
        }
        found = false
        for _,p:= range mypath{
            if v.x == p.x && v.y==p.y {
                found = true
                break
            }
        }

        dist:=abs(v.x-pt.x)+abs(v.y-pt.y)
        fmt.Fprintln(os.Stderr, dist, v, pt)
        if dist < minDist && dist >0 && !found{
            // fmt.Fprintln(os.Stderr,"min dist >>>>", dist, v, pt)

            minDist = dist
            nearestPt = v
        }

        if !found {
            newpaht = append(newpaht, v)
        }
    }
    // fmt.Fprintln(os.Stderr,"newpath: >>>>", newpaht)

    if onpath {
        fmt.Fprintln(os.Stderr,"onpath >>>>", nearestPt)
        return nearestPt
    }else{
        minDist = 100
        for _,v := range newpaht{
            dist:=abs(v.x-pt.x)+abs(v.y-pt.y)
            if dist < minDist {
                // fmt.Fprintln(os.Stderr,"min dist >>>>", dist, v, pt)
                minDist = dist
                nearestPt = v
            }
        }
        return nearestPt
    }

    // return newpaht[0]
}
func abs(value int) int{
    if value<0 {
        return value*-1
    }else{
        return value
    }
}
func scanMap(mypath[]cord)Data{
    var data Data
    for i := 0; i < 20; i++ {
            // line: One line of the map ('.' = free, '0' = you, otherwise the id of the opponent)
            var line string
            fmt.Scan(&line)
            var value int
            for idx, d :=range line{
                if d== '.' {
                    value = 0
                }else if d=='0' {
                    var found bool
                    for _,v:=range(mypath){
                        if v.x==idx && v.y==i{
                            found=true
                        }
                    }
                    if found{
                        value = 0
                    }else{
                        value = 9
                    }
                }else{
                    value = int(d -'0')
                }
                data[i][idx] = value
            }
            // fmt.Fprintln(os.Stderr, line)
        }
        return data
}
func print(mapState Data) {
    fmt.Fprintln(os.Stderr,"")
    for i := 0; i < 20; i++ {
        for j := 0; j < 35; j++ {
            fmt.Fprint(os.Stderr, mapState[i][j])
        }
        fmt.Fprintln(os.Stderr,"")
    }
}

func getRectList(mapState Data) []Rect{
    // print(mapState)
    transData := transform(mapState)
    // print(transData)
    rects := process(transData)
    return rects
}

func transform(mapState Data) Data{
    var data Data
    for i := 0; i < 20; i++ {
        for j := 0; j < 35; j++ {
            if i ==0 {
                if mapState[i][j] == 0 {
                    data[i][j] = 1
                }else{
                    data[i][j] = 0
                }
            }else{
                if mapState[i][j] == 0 {
                    data[i][j] = 1 + data[i-1][j]
                }else{
                    data[i][j] = 0
                }
            }
        }
    }
    return data
}
func process(tData Data) []Rect{
    var r Rect
    rectList := make([]Rect, 0, 0)
    for i := 19; i >= 0; i-- {
        r = largestArea(tData[i],i)

        // if r != nil {
            rectList = append(rectList, r)
        // }
    }
    return rectList
}

type Node struct {
    Value int
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
    nodes   []*Node
    count   int
}

func (s *Stack) push(n int) {
    s.Push(&Node{n})
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
    if s.count >= len(s.nodes) {
        nodes := make([]*Node, len(s.nodes)*2)
        copy(nodes, s.nodes)
        s.nodes = nodes
    }
    s.nodes[s.count] = n
    s.count++
}

func (s *Stack) pop() int {
    N:=s.Pop()
    if N!=nil {
        return N.Value
    }
    return -1
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
    if s.count == 0 {
        return nil
    }
    node := s.nodes[s.count-1]
    s.count--
    return node
}

func (s *Stack)empty() bool {
    return s.count==0
}
func (s * Stack)top() int {
    if s.count == 0 {
        return -1
    }
    node := s.nodes[s.count-1]
    return node.Value
}

func largestArea( arr[35] int, y int) Rect{

    // int area[len]; //initialize it to 0
    len:=35
    var  t int
    // stack<int> St;  //include stack for using this #include<stack>
    // var done bool
    var area[35] int
    var liIndex[35] int
    St := &Stack{nodes: make([]*Node, 35)}

    for i := 0; i < len; i++  {
        for !St.empty() {
           if arr[i] <= arr[St.top()] {
               St.pop()
           }else{
               break
           }
        }
        if St.empty(){
           t = -1
        }else{
           t = St.top()
        }
        //Calculating Li
        area[i] = i - t - 1
        liIndex[i] = area[i]
        St.push(i)
    }

    //clearing stack for finding Ri
    for !St.empty(){
        St.pop()
    }

    for i:=len-1; i>=0; i-- {
        for !St.empty(){
           if arr[i] <= arr[St.top()]{
               St.pop()
           }else{
               break
           }
        }
        if St.empty(){
           t = len;
        }else{
           t = St.top();
        }

        //calculating Ri, after this step area[i] = Li + Ri
        area[i] += t - i -1
        // fmt.Fprintln(os.Stderr,i,area[i] + 1,arr[i],t)
        St.push(i)
    }

    max := 0
    idx := 0
    var areaTmp[35] int

    //Calculating Area[i] and find max Area
    for i:=0; i<len; i++{
        // fmt.Fprintln(os.Stderr,area[i] + 1,arr[i])
        areaTmp[i] = arr[i] * (area[i] + 1);
        if areaTmp[i] > max{
           max = areaTmp[i];
           idx = i
        }
    }
    // fmt.Fprintln(os.Stderr,idx-liIndex[idx], y, area[idx]+1, arr[idx], max, liIndex[idx])
    // fmt.Fprintln(os.Stderr,"(",x-arr[idx]+1,idx,")","(",x,idx+area[idx],")", max, (area[idx] + 1), arr[idx])
    var r Rect
    // r.ltpt.x = x- liIndex[idx]
    r.ltpt.x = idx-liIndex[idx]
    r.ltpt.y = y-arr[idx]+1

    r.rbpt.x = r.ltpt.x + area[idx]
    r.rbpt.y = y
    r.area = max
    // fmt.Fprintln(os.Stderr,r)
    return r;
}

func getTop(rectList []Rect, pt cord) Rect{
    // maxArea := rectList[0].area
    // nearestDist := 100
    topIndex := 0
    minWeight := 100
    // nearestRank := 0
    for k,r := range rectList {
        // fmt.Fprintln(os.Stderr,k, r)
        if minWeight < r.area - getWeight(r,pt)*2 {
            minWeight = r.area - getWeight(r,pt)*2
            topIndex = k
        }
    }
    return rectList[topIndex]
}

func getWeight(rect Rect, pt cord) int{
    return 0
}
