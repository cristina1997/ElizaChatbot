// Creates constants for #list and #form id attributes from index
const list = $("#list");
const form = $("#form");

form.keypress(function(event){

    if(event.keyCode = 13){
         // form default refreshes the page and clears everything so we prevent it to make requests with JS
        event.preventDefault();

        const text = form.val(); // store the value taken from the form input into text
        form.val(""); // delete the value from the actual form input

        // if text is valid execute the following
        if (!text.trim()){
            return
        }
        
        list.append("<li>" + text + "</li>");

        $.get("/chat", {input: text})
        .done(function(w){ // "w" is the response from server
            list.append("<li>" + w + "</li>");
        }).fail(function(){ // it runs in case of an error (i.e. if the server isn't running)
            list.append("<li>Sorry! Eliza is not here. Please come back later!</li>");
        });
    }
    else if (event.keyCode != 13){
        return
    }
});