class Ports {
    constructor(container) {
        this.ports = container.ports
    }

    containerPorts() {
        var ports = this.ports;
        var port_numbers = new Array();

        for (var item in ports) {
            port_numbers.push(ports[item]["containerPort"]);
        }

        return port_numbers;
    }
}
