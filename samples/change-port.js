$$.forEach(function($) {
    var web_container = new Containers($.spec.template.spec.containers).by_name("web");
    web_container.Ports = [{ containerPort: 81 }, { containerPort: 444}];
});
