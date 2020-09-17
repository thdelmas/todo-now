# todo-now

## Table of Contents

* [About ToDo-Now](#about-todo-now)
  * [Roadmap](#roadmap)
  * [Built With](#built-with)
  * [Security](#security)
    * [Known Bugs](#known-bugs)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
  * [Testing](#testing)
  * [Usage](#usage)
* [API](#api)
  * [Data Architecture](#data-architecture)
* [Contributing](#contributing)
* [Contact Us](#contact-us)
* [Authors](#authors)
* [Licence](#licence)
* [Acknowledgments](#acknowledgments)
* [References](#references)

## About ToDo-Now
Todo list sort by optimization

### Roadmap

See the [open issues](https://github.com/thdelmas/todo-now/issues) for a list of proposed features.

### Built-With

### Security

#### Known Bugs

See the [open issues](https://github.com/thdelmas/todo-now/issues) for a list of known issues.

## Getting Started

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

```sh

```

### Installation

```sh

```

### Testing

How does someone test the code ?

```sh

```

### Usage

How does someone use the code ?

```sh

```

## API

## Data Architecture

### Event
Name | Required | Data Type
-----|----------|----------
name | True | str
owner | False | User
organizers | False | User list
attendees | False | User list
pretendants | False | User list
child events | False | Event list
topos | False | Topos list
constraints | False | constraint list
human_action | False | str
machine_action | False | shell script
descritpion | False | text
status | False | str


### Constraint
Name | Required | Data Type
-----|----------|----------
name | True | str
pipeline_complete_test | True | shell command, docker, CI, ?

### Topos
Name | Required | Data Type
-----|----------|----------
name | True | str
period | False | Period
area | False | Area

### Period
Name | Required | Data Type
-----|----------|----------
start_date | False | date
end_date |  False | date
timezone | True | ?

### Area
Name | Required | Data Type
-----|----------|----------
name | False | str
points | True | Point list

### Point
Name | Required | Data Type
-----|----------|----------
altitude | True | float
longitude | True | float
latitude | True | float

## Contributing

How does someone develop the code ?

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/).

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/NewFeature`)
3. Commit your Changes (`git commit -m 'Add some fun here'`)
4. Push to the Branch (`git push origin feature/NewFeature`)
5. [Open a Pull Request](https://github.com/thdelmas/todo-now/compare/).

## Contact Us

## Authors

* **Théophile Delmas** - *Initial design*

See also the list of [contributors](https://github.com/thdelmas/todo-now/graphs/contributors) who participated in this project.


## Licence

License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

## References

[ntask](https://www.ntaskmanager.com/)

[wunderlist](https://www.wunderlist.com/)

[Todoist](https://todoist.com/)

[Habitica](https://habitica.com/)

[Remember the milk](https://www.rememberthemilk.com/)
