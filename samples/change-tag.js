$$.forEach(function($) {
    var containers = new Containers($.spec.template.spec.containers);
    var web_container = containers.by_name("web");
    var containerImage = new DockerImage(web_container.image);
    containerImage.tag = "1.2";
    web_container.image = containerImage.address();
});
