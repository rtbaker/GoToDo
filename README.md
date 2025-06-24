# GoToDo

Simple ToDo API. 

A Todo has a title, description, priority and a completed flag.

## Structure

### Model package

Defines an interface for the data we want to persist, doing it this way means we can
implement this interface for each database type we want to support.

### Misc

cmd/gogotod/main -> 
    main entry point -> 
    creates application that sorts out config -> 
    creates http server that sets up listening and rounting