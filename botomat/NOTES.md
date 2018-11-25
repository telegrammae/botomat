# About the challenge
This challenge was a nice opportunity to use the new Go modules introduced in Go 1.11.
The resulting solution has a module called 'botomat', a command-line program.

## Features
On top of the core functionality, botomat also has two additional features: the ability to run multiple Robots at the same time and the significance of a Robot's type.

## Design
The main design decision revolved around entities and their methods, as well as access to a shared resource. Since a Robot deletes a task upon finishing it, that means access to the global list of tasks must be shared among Robots.

In general, using a slice is not so easy, since removing elements from a slice is problematic because of multiple goroutines accessing it and modifying it.

Using some kind of map is more convenient. Also, it allows for further control over the state.
If the keys of a map are tasks themselves, the values can be boolean values, which indicate if a task has been acquired by a Robot. By default, all values are false.

However, using a plain map is not possible, since concurrent modifications to it are nearly unavoidable. Another option was to use a map with a mutex guarding access to it during task deletion.
But there still would be problems with concurrent access because of the concurrent loops that occur when multiple Robots spawn multiple goroutines to work on the tasks.

Another solution is to try and use a concurrent map, sync.Map, introduced in Go 1.9.
Along with the mechanism to control the precise moment when the value of the map's entry is set to true, it is possible to have multiple Robots work only on distinct tasks and avoid concurrent writes and any inconsistent states.

In general, users are free to create many Robots, with each being able to spawn a different number of goroutines. Each goroutines processes a task.
Of course, there is a trade-off in terms of performance, depending on the exact configuration.
For example, it may be better to use two Robots, with each creating 10 goroutines, instead of four Robots, with each creating two goroutines.

As for the object-oriented design of this program, I am not completely happy with it still.
Overall, it makes sense to keep the Robot implementation package-private and only expose a few methods of the factory, the Botomat itself.
The small issue seems to be that Robots receive a pointer to the map from the factory and thus modify it, as necessary. As a result, the Robot and factory are somewhat strongly coupled.
It could, however, be easily changed, should the design expand, but only at this early point.
Botomat returns an instance of Robot and allows to run it on a collection of tasks.
A Robot can modify the collection of tasks of its Botomat. A Robot can be created and run only in the context of its Botomat.

I made a function to create a larger collection of tasks for easier testing of the program. The tasks themselves are equivalent to those listed in the README.

One requirement, from which I deviated slightly, was to assign a robot five tasks. I made it so that each robot works on as many tasks at a time as its type (model) specifies. I felt that the challenge was open-ended enough to allow for this deviation.
