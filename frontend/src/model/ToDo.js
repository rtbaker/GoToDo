// Utility class to take the JSON from the api into something more usable

export default class ToDo {
    constructor(
        id,
        title,
        description,
        updatedAt,
        priority,
        completed,
    ) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.updatedAt = updatedAt;
        this.priority = priority;
        this.completed = completed;
    }

    static newFromJSON(data) {
        const updatedAt = new Date(data.updatedAt);
        return new ToDo(data.id, data.title, data.description, updatedAt, data.priority, data.completed);
    }

    displayDate() {
        return this.updatedAt.toLocaleDateString();
    }
}
