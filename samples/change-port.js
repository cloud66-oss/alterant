$$.forEach($ => {
    var web_container = new Containers($.spec.template.spec.containers).by_name("web");
    var ports = [{ containerPort: 81 }, { containerPort: 444}];
    web_container.ports = ports;
});
