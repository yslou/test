function initMap() {
    // Create a map object and specify the DOM element for display.
    var map = new google.maps.Map(document.getElementById('map'), {
        center: {lat: 25.0329636, lng: 121.5654268},
        scrollwheel: false,
        zoom: 18
    });

    var marker = new google.maps.Marker({
        position: {lat: 25.0329636, lng: 121.5654268},
        map: map,
        label: 'C',
        title: 'Hello World!'
    });

    var marker2 = new google.maps.Marker({
        position: {lat: 25.0339636, lng: 121.5654268},
        map: map,
        label: 'A chan',
        title: 'Hello yyy!',
        icon: { path:google.maps.SymbolPath.BACKWARD_OPEN_ARROW, scale:10} 
    });
}

 

