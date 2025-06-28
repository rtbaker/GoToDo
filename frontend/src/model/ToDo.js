// Utility class to take the JSON from the api into something more usable

export default class ToDo {
    constructor(
        id,
        title,
        description,
        updatedAt
    ) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.updatedAt = updatedAt;
    }

    static newFromJSON(data) {
        const updatedAt = new Date(data.updatedAt);
        return new ToDo(data.id, data.title, data.description, updatedAt);
    }

    displayDate() {
        return this.updatedAt.toLocaleDateString();
    }
}
