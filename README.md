# Bellina API Reference

Bellina is a low-level framework for building flexible graphical user interfaces in GoLang.

Refer to this document for the specification of the public API for Bellina.

High-level explanations, Tutorials, Setups, and Examples can be found at the following links:

 * [Bellina Public API Specification (Github)](https://github.com/amortaza/go-bellina)
 * [Bellina Tutorials (Github)](https://github.com/amortaza/go-bellina-tutorials)
 
&nbsp;

Some of the concepts referred to in this document (e.g. current node context) will not make sense without an introduction to the Bellina model.  Refer to the links above for basic introductions and tutorials.

&nbsp;

## Public API
<div id="toc-node" />
### Current Node Context API

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 
[Node Struct](#node-struct)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 
[Shadow Node Struct](#shadownode-struct)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 
[bl.Root()](#blroot)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 
[bl.Div()](#bldiv)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Id(id string)](#blid)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.End()](#blend)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Pos(left, top int)](#blpos)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Dim(width, height int)](#bldim)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.InvisibleToMouseEvents()](#blinvisibletomouseevents)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.CustomRenderer(f func(node *Node), topsKids bool))](#blcustomrenderer)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Dirty()](#bldirty)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.SettleBoundary()](#blsettleboundary)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.SettleKids()](#blsettlekids)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.RequireSettledBoundary()](#blrequiresettledboundary)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.RequireSettledKids()](#blrequiresettledkids)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OwnLeft(owner string) bool](#blownltwh)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OwnTop(owner string) bool](#blownltwh)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OwnWidth(owner string) bool](#blownltwh)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OwnHeight(owner string) bool](#blownltwh)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OwnPos(owner string) bool](#blownpos)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OwnDim(owner string) bool](#blowndim)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OnMouseMove(cb func(*MouseMoveEvent))](#blonmousemove)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.OnMouseButton(cb func(*MouseButtonEvent))](#blonmousebutton)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.AddFunc(cb func())](#bladdfunc)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.EnsureShadow() *ShadowNode](#blensureshadow)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.EnsureShadowById(id string) *ShadowNode](#blensureshadowbyid)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.EnsureShadowByNode(node *Node) *ShadowNode](#blensureshadowbynode)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.HasShadowById(id string) (*ShadowNode, bool)](#blhasshadowbyid)

<div id="toc-events" />
### Events API

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Event interface](#blevent)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.RegisterLongTerm(eventType string, callback func(Event))](#blregisterlongterm)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.RegisterShortTerm(eventType string, callback func(Event))](#blregistershortterm)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.FireEvent(event Event)](#blfireevent)

<div id="toc-lc" />
### Life-Cycle API

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Register_LifeCycle_After_UserTick_ShortTerm(cb func())](#blregisterlifecycleafterusertickshortterm)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Register_LifeCycle_After_UserTick_LongTerm(cb func())](#blregisterlifecycleafteruserticklongterm)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Register_LifeCycle_Before_UserTick_LongTerm(cb func())](#blregisterlifecyclebeforeuserticklongterm)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Register_LifeCycle_Init(cb func())](#blregisterlifecycleinit)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Register_LifeCycle_Uninit(cb func())](#blregisterlifecycleuninit)

<div id="toc-helper" />
### Helper API

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.GetNodeAbsolutePos(node *Node)(absX, absY int)](#blgetnodeabsolutepos)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.DivId(id string)](#bldivid)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.GetNodeById(id string ) *Node](#blgetnodebyid)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Stabilize(node *Node)](#blstabilize)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Disp(node *Node)](#bldisp)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.SetMouseCursor(cursor MouseCursor)](#blsetmousecursor)

<div id="toc-hal" />
### HAL API (Hardware Abstraction Layer)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[Canvas Interface](#icanvas)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[Graphics Interface](#igraphics)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[HAL Interface](#ihal)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Start(...)](#blstart)

<div id="toc-globals" />
### Global Variables

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Root_Node](#blrootnode)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Current_Node](#blcurrentnode)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Window_Width, bl.Window_Height](#blwindowwh)

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
[bl.Mouse_X, bl.Mouse_Y](#blmousexy)

&nbsp;

## Current Node Context API

<div id="node-struct" />
```
type Node struct {

	Id                       string

	Left, Top, Width, Height int

	Parent                   *Node
	Kids                     *list.List

	OnMouseMoveCallbacks     *list.List
	OnMouseButtonCallbacks   *list.List

	CustomRender_1           func(node *Node)
	CustomRender_2           func(node *Node)
	CustomRenderTopsKids     bool

	InvisibleToMouseEvents   bool

	OwnerOfLeft              string
	OwnerOfTop               string
	OwnerOfWidth             string
	OwnerOfHeight            string

	SettledBoundary          bool
	SettledKids              bool

	Dirty                    bool
}
```

<div id="shadownode-struct" />
```
type ShadowNode struct {
	Id                       string

	Left, Top, Width, Height int

	BackingNode              *Node // this is the node that the shadow is backing!!
}
```

&nbsp;

<div id="blroot" />
#### Definition

```
bl.Root()
```

- Starts a new Bellina node tree
- This is done in the *User Tick callback function* (see Bellina Programming PDF for details on the life-cycle design)
- The *root* node has ID of "ROOT"
- After call to `bl.Root()`:
   - The global variable `bl.Root_Node` points to the root node of the Bellina tree
   - The global variable `bl.Current_Node` also points to the root node (for the moment)
   - The *node-cache* can return the root node via call `bl.GetNodeById("ROOT")`
   - A context has begun for the new node and further calls affect this new node
   - The context *must* be closed with call to `bl.End()` when it is no longer needed

#### Example Usage

```
bl.Root()
{
    bl.Pos(50, 75)  // position the node
    bl.Dim(60, 40)  // size the node
}
bl.End()
```

`bl.Root()` is almost identical to

```
bl.Div()
{
   bl.Id("ROOT")
   ...
```

The difference is `bl.Root()` sets the global variable `bl.Root_Node` to the new node.

&nbsp;

---
<div id="bldiv" />
#### Definition

```
bl.Div()
```

- Starts a new node within the current context (a child-node)
- After call to `bl.Div()`:
   - The global variable `bl.Current_Node` points to the newly created node
   - A context has begun for the new node and further calls affect this new node
   - The `Parent` property of the new node is set
   - The context *must* be closed with call to `bl.End()` when it is no longer needed

**Note** The node will not have any properties (except `Parent`) set (e.g. Id, position, dimension)

Consequently the *node-cache* can *NOT* return the node via call `bl.GetNodeById(id string)` until `bl.Id(id string)` is called.

`bl.Div()` can be called in an arbitrarily nested fashion.

#### Example Usage

```
bl.Root()
{
    bl.Pos(50, 75)  // position the node
    bl.Dim(60, 40)  // size the node

    bl.Div()
    {
    	bl.Id("child")
        bl.Pos(10, 20)
        bl.Dim(60, 90)

        bl.Div()
        {
            bl.Id("grand-child")
            bl.Pos(10, 20)
            bl.Dim(60, 90)
        }
        bl.End()
    }
    bl.End()
}
bl.End()
```

&nbsp;

---
<div id="blid" />
#### Definition

```
bl.Id(id string)
```

- Sets the ID of the current context node (i.e. `bl.Current_Node`)
- After call to `bl.Id(id string)`:
   - The global variable `bl.Current_Node` ID will be `id`
   - The *node-cache* will return the node via call `bl.GetNodeById(id string)`

#### Notes
- `bl.Root()` does not require call to `bl.Id(id string)` because the root node will automatically get value 'ROOT' as the ID


- An ID is required for every node in the tree

#### Example Usage

```
bl.Root()
{
    bl.Pos(50, 75)  // position the node
    bl.Dim(60, 40)  // size the node

    bl.Div()
    {
    	bl.Id("child")
        bl.Pos(10, 20)
        bl.Dim(60, 90)
    }
    bl.End()
}
bl.End()
```

&nbsp;

---
<div id="blend" />
#### Definition

```
bl.End()
```

- Closes the current node context
- After call to `bl.End()`:
   - The global variable `bl.Current_Node` points to the parent of the previously active context node

**Note** At the `bl.Root()` level, the `bl.End()` sets `bl.Current_Node` to `nil` since we have effectively exited the tree.

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        // at this point bl.Current_Node is 'child' node
    }
    bl.End()

    // at this point bl.Current_Node is 'root' node
}
bl.End()

// at this point bl.Current_Node is nil

```

&nbsp;

---
<div id="blpos" />
#### Definition

```
bl.Pos(left, top int)
```

- Sets the `Left` and `Right` properties of the current node
- The coordinates are relative to the node's parent's (top, left)
- A position of (0,0), puts the node at the node's parent's top-left corner

#### Example Usage

```
bl.Root()
{
    // sets the Left, Top properties of 'ROOT' node
	bl.Pos(200, 300)

    bl.Div()
    {
    	bl.Id("child")

        // sets the Left, Top properties of 'child' node
        bl.Pos(100, 400)
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="bldim" />
#### Definition

```
bl.Dim(width, height int)
```

- Sets the `Width` and `Height` properties of the current node

#### Example Usage

```
bl.Root()
{
    // sets the Width, Height properties of 'ROOT' node
	bl.Dim(200, 300)

    bl.Div()
    {
    	bl.Id("child")

        // sets the Width, Height properties of 'child' node
        bl.Dim(100, 400)
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blinvisibletomouseevents" />
#### Definition

```
bl.InvisibleToMouseEvents()
```

- Sets the `InvisibleToMouseEvents` property of the current node to true
- While this property is `true`, any mouse move or mouse button events that would have otherwise been triggered on this node, are ignored
- If a node is invisible to mouse events, so are all of its child nodes

**Note** By default, this property is `false`.  And since the node tree is re-generated every frame, `bl.InvisibleToMouseEvents()` is the only way to set this property to `true`.

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        // set the current nodes to be invisible to mouse events
        bl.InvisibleToMouseEvents()
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blcustomrenderer" />
#### Definition

```
bl.CustomRenderer(f func(node *Node), topsKids bool)
```

- Sets the `CustomRenderer` property of the current node to function `f`
- When the node is about to be rendered, the custom renderer function is called to do the actual rendering
- If `topsKids` is true, the child nodes are rendered first, then the node is rendered on top
- If `topsKids` is false, the node is rendered first, then the child nodes are rendered on top

A node can have 2 custom renderers which are invoked 1 then 2 (i.e. 2nd renderer draws on top of the 1st).

The first call to `bl.CustomRenderer(...)` sets the 1st renderer.

The second call TO `bl.CustomRenderer(...)` sets the 2nd renderer.

If a node does not have a customer renderer function set, then nothing is drawn - *however* the child nodes are traversed and rendered unobstructed.

A node's custom renderer function settings does not affect the node's ability to receives events.  The two abilities do not interfere with one another.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        bl.CustomRenderer(func(node *bl.Node), true) {
            // do rendering functions given the 'node' is passed in.
        })
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="bldirty" />
#### Definition

```
bl.Dirty()
```

- Sets the `Dirty` property of the current node to `true`

When a node is rendered, its image is cached for reuse.  As long as the node is not `Dirty`, the cached image is reused for rendering at great speed-up.

Bellina cannot determine if a node is `Dirty` and needs to have its image refreshed, since the custom renderer function for the node can use properties that are not Bellina's scope of understanding.

When a developer determines that a node is `Dirty`, they can call the `bl.Dirty()` function and flag the node to be redrawn.

**Note** When a node is `Dirty`, the chain of its parents are also marked as `Dirty`.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        if stateChanged {
            bl.Dirty()
        }
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blsettleboundary" />
#### Definition

```
bl.SettleBoundary()
```

- Sets the flag on a node that indicates its dimension and position (aka boundary) has been set or settled

This comes into play when, say, a layout plugin needs to ensure that the node's boundaries have been settled so it can be sure that is calculations are accurate.  For example, if a layout makes assumptions, only to have those assumptions be overwritten, the node will have the wrong dimension or position.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")
        bl.Pos(10, 10)
        bl.Dim(20, 20)

        // signal others that the position and dimention (aka boundary)
        // of this node has been set and will not be set again
        bl.SettleBoundary()

        // now, say, a layout plugin can assert that the boundary will not change as it does its calculations
        ...
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blsettlekids" />
#### Definition

```
bl.SettleKids()
```

- Sets the flag on a node that indicates its done with adding child nodes

This comes into play when, say, a layout plugin needs to ensure that the node's kids have been settled so it can be sure that is calculations are accurate.  For example, if a layout makes assumptions, only to have those assumptions be overwritten, the node will have the wrong dimension or position.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        // create kids
        bl.Div()
        {
        	// ...
        }
        bl.End()

        // signal that no more kids will be added
        bl.SettleKids()

        // now, say, a layout plugin can assert that the number of
        // child nodes will not change as it does its calculations
        ...
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blrequiresettledboundary" />
#### Definition

```
bl.RequireSettledBoundary()
```

- Assert that `bl.SettleBoundary()` has been called on this node

See `bl.SettleBoundary()` definition

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        bl.Pos(10,10)
        bl.Dim(50,50)

        // signal that position nor dimension will be modified
        bl.SettleBoundary()

        // now, say, a layout plugin can assert that the dimension or position of node will not change
        ...

        // plugin can validate current node has settled boundary
        bl.RequireSettledBoundary()
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blrequiresettledkids" />
#### Definition

```
bl.RequireSettledKids()
```

- Assert that `bl.SettleKids()` has been called on this node

See `bl.SettleKids()` definition

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        // create kids
        bl.Div()
        {
        	// ...
        }
        bl.End()

        // signal that no more kids will be added
        bl.SettleKids()

        // now, say, a layout plugin can assert that the number of
        // child nodes will not change as it does its calculations
        ...

        // plugin can validate current node has settled kids
        bl.RequireSettledKids()
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blownltwh" />
#### Definition

```
bl.OwnLeft(newOwner string) bool
bl.OwnTop(newOwner string) bool
bl.OwnWidth(newOwner string) bool
bl.OwnHeight(newOwner string) bool
```

All *Own* *left*, *top*, *width*, *height* functions work similarly.

- *Attempts* to set the `OwnerOfLeft` of the current node to `newOwner`
- returns `true` if successful

**Note** `newOwner` value of `'*'` is a special owner that overwrites other owners.
See the rules below.

### Case when existing owner is `''` or `'*'`

- The `OwnerOfLeft` of the node will be set to `newOwner`
- `true` will be returned for successful

### Case when existing owner is same as `newOwner`
- `true` will be returned for successful

### Case when new owner is `'*'`

- Regardless of existing *owner*, node's `OwnerOfLeft` is not modified
- *However*, `true` is returned for successful

### Other remaining case

- When `newOwner` is not the same as the existing `OwnerOfLeft`, then an warning is printed
- `false` is returned for unsuccessful

### Proper usage
These *Own* functions are mostly called within Plugins.  If a plugin needs to modify the Left property of a node, and another plugin modifies the same node's Left property, unexpected results can happen.

Using the *Own* functions to ensure correct ownership will reduce plugins stepping on each other's results.

When a plugin is overwritting another plugin's results and this overwritting is expected, then use the `'*'` owner to overwrite ownership and remove warning messages.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        if bl.OwnLeft("me") {
            // "me" now owns 'Left' property of node
        }
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blownpos" />
#### Definition

```
bl.OwnPos(newOwner string) bool
```

- Calls `bl.OwnLeft(...)` and `bl.OwnTop(...)`
- returns `true` if both calls return `true`

**Note** **Both** functions are called, regardless of their return values.

See `bl.OwnLeft(...)` and `bl.OwnTop(...)` for further details.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        if bl.OwnPos("me") {
            // "me" now owns 'Left' and 'Top' properties of node
        }
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="bldim" />
#### Definition

```
bl.OwnDim(newOwner string) bool
```

- Calls `bl.OwnWidth(...)` and `bl.OwnHeight(...)`
- returns `true` if both calls return `true`

**Note** **Both** functions are called, regardless of their return values.

See `bl.OwnWidth(...)` and `bl.OwnHeight(...)` for further details.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        if bl.OwnDim("me") {
            // "me" now owns 'Width' and 'Height' properties of node
        }
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blonmousemove" />
#### Definition

```
bl.OnMouseMove(cb func(*MouseMoveEvent))
```

- Registers `cb` callback function with a *mouse move* event for the current node
- `OnMouseMove(...)` may be called many times to register different callback functions

Since node tree is re-created every frame, the list of registered callbacks needs to be rebuilt every frame.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

See structures section in this document for definition of `bl.MouseMoveEvent`

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        bl.OnMouseMove(func(event *bl.MouseMoveEvent) {
            fmt.Printf("Mouse was moved over node ", event.Target.Id)
        })
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blonmousebutton" />
#### Definition

```
bl.OnMouseButton(cb func(*MouseButtonEvent))
```

- Registers `cb` callback function with a *mouse button* event for the current node
- `OnMouseButton(...)` may be called many times to register different callback functions

Since node tree is re-created every frame, the list of registered callbacks needs to be rebuilt every frame.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

See structures section in this document for definition of `bl.MouseButtonEvent`

#### Example Usage

```
bl.Root()
{
    bl.Div()
    {
    	bl.Id("child")

        bl.OnMouseButton(func(event *bl.MouseButtonEvent) {
            fmt.Printf("Mouse button pressed over node ", event.Target.Id)
        })
    }
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="bladdfunc" />
#### Definition

```
bl.AddFunc(cb func())
```

- `bl.AddFunc(...)` may be called any number of times for any given node


- Registers `cb` callback function to be called during the `Stabilization` portion of the Bellina loop life-cycle


- See `Life-Cycle` section of [Bellina Programming Guide](todo)


- Any `bl.AddFunc(...)` called *before* child nodes have been created, will get the registered `cb` function called before any of the child nodes' registered callback functions are called


- Any `bl.AddFunc(...)` called *after* child nodes have been created, will get the registered `cb` function called after the child nodes' registered callback functions are called


- These `cb` are called in order of:
    - callbacks that were registered *before* child nodes were added
    - callbacks registered within child nodes (followed *recursively*)
    - callbacks that were registered *after* child nodes were added


Since node tree is re-created every frame, the list of registered callbacks needs to be rebuilt every frame.

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

See `Life-Cycle` section of [Bellina Programming Guide](todo)

#### Example Usage

The following example prints:

```
Hello Parent
Hello Child
Bye Parent
```

```
bl.Root()
{
    bl.Div()
    {
        bl.Id("A")

        bl.AddFunc(func() {
            fmt.Printf("Hello Parent")
        })

        bl.Div()
        {
            bl.Id("B")

	        bl.AddFunc(func() {
    	        fmt.Printf("Hello Child")
        	})
        }
        bl.End()

        bl.AddFunc(func() {
            fmt.Printf("Bye Parent")
        })
}
    bl.End()
}
bl.End()

```

&nbsp;

---
<div id="blensureshadow" />
#### Definition

```
bl.EnsureShadow() *ShadowNode
```

- Returns the `ShadowNode` for the current node, or creates one if it does not exist


- `ShadowNode`s persist across frames and always maintain their properties


- `ShadowNode`'s `BackingNode` property always points to the latest copy of the `Node` with the same `Id` (keep in mind that every frame, all nodes are destroyed


&nbsp;

---
<div id="blensureshadowbyid" />
#### Definition

```
bl.EnsureShadowById(id string) *ShadowNode
```

- Returns the `ShadowNode` for the node with given `id`, or creates one if it does not exist


- `ShadowNode`s persist across frames and always maintain their properties


- `ShadowNode`'s `BackingNode` property always points to the latest copy of the `Node` with the same `Id` (keep in mind that every frame, all nodes are destroyed

&nbsp;

---
<div id="blensureshadowbynode" />
#### Definition

```
bl.EnsureShadowByNode(node *Node) *ShadowNode
```

- Returns the `ShadowNode` for the node or creates one if it does not exist


- `ShadowNode`s persist across frames and always maintain their properties


- `ShadowNode`'s `BackingNode` property always points to the latest copy of the `Node` with the same `Id` (keep in mind that every frame, all nodes are destroyed

&nbsp;

---
<div id="blhasshadowbyid" />
#### Definition

```
bl.HasShadowById(id string) (*ShadowNode, bool)
```

- Returns the `ShadowNode` for the node if it exists


- Returns `false` if the `ShadowNode`the `ShadowNode` for the node if it exists


&nbsp;

&nbsp;


## Events API

```
type Event interface {
	Type() string
}

type KeyEvent struct {
	Key KeyboardKey
	Action ButtonAction
	Alt, Ctrl, Shift bool
}

type MouseButtonEvent struct {
	Button         MouseButton
	ButtonAction   ButtonAction
	Target         *Node
}

type MouseMoveEvent struct {
	X, Y int
	Target *Node
}

type WindowResizeEvent struct {
	Width, Height int
}

```

&nbsp;

---
<div id="blregisterlongterm" />
#### Definition

```
bl.RegisterLongTerm(eventType string, cb func(Event))
```

- Registers a `cb` in the *long-term* callback queue
- *long-term* queue does not get cleared by Bellina
- Many callbacks may be registered for a given event type
- Only events that are fired via `bl.FireEvent(event Event)` respect this registration

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Notes

- `bl.RegisterLongTerm(...)` is called usually during *initialization* phase of a plugin
- `bl.RegisterLongTerm(...)` should not be called every frame as the queue will keep getting bigger with duplicates

#### Example Usage

```
// a hover plugin might register long term listener for the mouse move event

bl.RegisterLongTerm( bl.EventType_Mouse_Move, func(e bl.Event) {
		event := e.(*bl.MouseMoveEvent)

        // wherever a mouse is moved globally, this function will get called
        // it is not tied to a node - but the event will say which node this
        // even comes from
})
```

&nbsp;

---
<div id="blregistershortterm" />
#### Definition

```
bl.RegisterShortTerm(eventType string, cb func(Event))
```

- Registers a `cb` in the *short-term* callback queue
- *short-term* queue is cleared by Bellina every frame
- Many callbacks may be registered for a given event type
- Only events that are fired via `bl.FireEvent(event Event)` respect this registration

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
// a click plugin might register short term listener for the mouse button event

bl.RegisterShortTerm(bl.EventType_Mouse_Button, func(event bl.Event) {
		event := e.(*bl.MouseMoveEvent)

        // wherever a mouse is clicked on a node, this function gets called
})
```

#### Question

Could the plugin above has used a `RegisterLongTerm` ?

I think yes, but it is prefered to use the *short-term* whenever possible.  The idea is that all data structures are destroyed and recreated every frame.  So if something can be done with *short* term, this it should be.

Sometimes you have no choice but to use *long* term.  The tutorials cover these advanced subjects.


&nbsp;

---
<div id="blfireevent" />
#### Definition

```
FireEvent(event Event)
```

- Calls all `cb` in the *long-term* callback queue that match `event`'s type
- Calls all `cb` in the *short-term* callback queue that match `event`'s type

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Example Usage

```
    // Fire a key event, or even simulate a keypress!

	keyEvent := NewKeyEvent(key, action, alt, ctrl, shift)

	FireEvent(keyEvent)
```


&nbsp;

&nbsp;


## Life-Cycle API
<div id="blregisterlifecycleafterusertickshortterm" />
#### Definition

```
bl.Register_LifeCycle_After_UserTick_ShortTerm(cb func())
```

- Registers a `cb` in the *short-term* *after* user tick callback queue
- `cb`s in this queue get called after the `user-tick` function
- This queue gets cleared every frame - hence the name *short-term*

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Use case

TODO - but see the tutorials

&nbsp;

---
<div id="blregisterlifecycleafteruserticklongterm" />
#### Definition

```
bl.Register_LifeCycle_After_UserTick_LongTerm(cb func())
```

- Registers a `cb` in the *long-term* *after* user tick callback queue
- `cb`s in this queue get called after the `user-tick` function
- This queue does not get cleared - hence the name *long-term*

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Use case

TODO - but see the tutorials

&nbsp;

---
<div id="blregisterlifecyclebeforeusertickshortterm" />
#### Definition

```
bl.Register_LifeCycle_Before_UserTick_ShortTerm(cb func())
```

- Registers a `cb` in the *short-term* *before* user tick callback queue
- `cb`s in this queue get called before the `user-tick` function
- This queue gets cleared every frame - hence the name *short-term*

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Use case

TODO - but see the tutorials

&nbsp;

---
<div id="blregisterlifecyclebeforeuserticklongterm" />
#### Definition

```
bl.Register_LifeCycle_Before_UserTick_LongTerm(cb func())
```

- Registers a `cb` in the *long-term* *before* user tick callback queue
- `cb`s in this queue get called before the `user-tick` function
- This queue does not get cleared - hence the name *long-term*

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Use case

TODO - but see the tutorials


&nbsp;

---
<div id="blregisterlifecycleinit" />
#### Definition

```
bl.Register_LifeCycle_Init(cb func())
```

- Registers a `cb` in the *one-time* called callback queue
- `cb`s in this queue get called once when the Bellina system has initialized

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Use case

TODO - but see the tutorials

&nbsp;

---
<div id="blregisterlifecycleuninit" />
#### Definition

```
bl.Register_LifeCycle_Uninit(cb func())
```

- Registers a `cb` in the *one-time* called callback queue
- `cb`s in this queue get called once when the Bellina system has uninitialized

For an in-depth example, see github.com/amortaza/go-bellina-tutorials

#### Use case

TODO - but see the tutorials

&nbsp;

&nbsp;


## Helper API
<div id="blgetnodeabsolutepos" />
#### Definition

```
bl.GetNodeAbsolutePos(node *Node)(absX, absY int)
```

- Returns the absolute coordinates of a node relative to the application window

&nbsp;

---
<div id="bldivid" />
#### Definition

```
bl.DivId(id string)
```

- Behaves like `bl.Div()` except does not create a new node, but takes the node identified by `id` and makes it the current context

&nbsp;

---
<div id="blgetnodebyid" />
#### Definition

```
bl.GetNodeById(id string ) *Node
```

- Returns the node identified by `id`

- If the node does not exist then an exception occurs in a form of panic

&nbsp;

---
<div id="bldisp" />
#### Definition

```
bl.Disp(node *Node)
```

- Prints information about the node (e.g. its properties)

&nbsp;

---
<div id="blsetmousecursor" />
#### Definition

```
bl.SetMouseCursor(cursor MouseCursor)
```

- Changes the mouse cursor

Valid values for `cursor` are

```
const (
	MouseCursor_Arrow
	MouseCursor_Horiz_Resize
	MouseCursor_Vert_Resize
	MouseCursor_IBeam
	MouseCursor_CrossHair
	MouseCursor_Hand
)
```

#### Example

```
bl.SetMouseCursor(bl.MouseCursor_IBeam)
```

&nbsp;

## HAL API (Hardware Abstraction Layer)

<div id="icanvas" />
```
type Canvas interface {
	Begin()
	End()
	GetWidth() int
	GetHeight() int
	Clear(red, green, blue float32)
	Paint(seeThru bool, left, top int, alphas []float32)
	Free()
}
```

<div id="igraphics" />
```
type Graphics interface {
	Clear(red, green, blue, alpha float32)
	PushView(width, height int)
	PopView()
	NewCanvas(width, height int) Canvas
}
```

<div id="ihal" />
```
type HAL interface {

	Start(	width, height int,
		title string,
		onAfterGL, onLoop, onBeforeDelete func(),
		onResize, onMouseMove func(int,int),
		onMouseButton func(MouseButton, ButtonAction),
		onKey func(KeyboardKey, ButtonAction, bool, bool, bool))

	GetWindowDim()(width, height int)

	GetMousePos()(x,y int)

	SetMouseCursor(cursor MouseCursor)

	GetGraphics() Graphics
}
```

<div id="blstart" />
#### Definition

```
bl.Start(	hal HAL,
			width, height int,
			title string,
			init func(),
			tick func(),
			uninit func())
```

- This is how to start a Bellina application

#### Example

A minimum Bellina application

```
package main

import (
	"runtime"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
)

func initialize() {
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(1024, 768)
	}
	bl.End()
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( hal_g5.New(), 1280, 1024, "Bellina v0.9", initialize, tick, nil )
}

```



&nbsp;

## Global Variables
<div id="blrootnode" />
`bl.Root_Node`

- Pointer to the root node of the Bellina tree
- Is set during call to `bl.Root()`
- Think of this tree as the DOM of a Bellina application

&nbsp;
<div id="blcurrentnode" />
`bl.Current_Node`

- While within a `bl.Div()` / `bl.End()` pair, `bl.Current_Node` points to the current active node
- All functions are operating on `bl.Current_Node`
- Is set during call to `bl.Root()`, `bl.Div()`, and `bl.DivId(id string)`

&nbsp;
<div id="blwindowwh" />
`bl.Window_Width`, `bl.Window_Height`

- The current *application* Window width and height
- Is always up to date

&nbsp;
<div id="blmousexy" />
`bl.Mouse_X`, `bl.Mouse_Y`

- The current mouse (x, y) coordinates relative to the *application* Window
- Is always up to date

