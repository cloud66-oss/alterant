class Containers {
    constructor(containers) {
        this.containers = containers
    }

    by_name(name) {
        for (var c in this.containers) {
            var item = this.containers[c]
            if (item.name == name) {
                return item
            }
        }

        return null
    }
}
