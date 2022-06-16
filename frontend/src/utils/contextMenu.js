export const contextMenu = (element, event) => {
    var contextElement = document.getElementById(element);
    contextElement.style.display = 'block';

    var screenWidth = window.screen.width;

    if( screenWidth - event.clientX < 200 ){
      contextElement.style.left = contextElement.style.top  = 0 + "px";
      contextElement.style.top = event.clientY + 11 + "px";
      contextElement.style.left = event.clientX - 232 + "px";
    }else{
      contextElement.style.left = contextElement.style.top = 0 + "px";
      contextElement.style.top = event.clientY + 11 + "px";
      contextElement.style.left = event.clientX  + 11 + "px";
    }
}