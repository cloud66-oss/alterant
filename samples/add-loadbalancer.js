$$.forEach($ => {
    if ($.kind == 'Service') {
        $.spec.type = "LoadBalancer";
        //$.spec.loadBalancerIP = "35.199.15.224";
    }
});
