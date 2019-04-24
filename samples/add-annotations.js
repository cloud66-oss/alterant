$$.forEach(function($){
    if ($.kind === 'Service') {
        console.log("got it")
        $.metadata.annotations = [{ "cloud66.com/deployed-at": Date.now() }];
    }
});
