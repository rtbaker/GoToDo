// Utility class to take the JSON from the api into something more usable

export default class ToDo {
    constructor(
        id,
        title,
        description,
        updatedAt,
        createdAt,
        priority,
        completed,
    ) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.updatedAt = updatedAt;
        this.createdAt = createdAt;
        this.priority = priority;
        this.completed = completed;
    }

    static newFromJSON(data) {
        const updatedAt = new Date(data.updatedAt);
        const createdAt = new Date(data.createdAt);
        return new ToDo(data.id, data.title, data.description, updatedAt, createdAt, data.priority, data.completed);
    }

    displayUpdatedDate() {
        return this.updatedAt.toLocaleDateString();
    }

    displayCreatedDate() {
        return this.createdAt.toLocaleDateString();
    }
}
