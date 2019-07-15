var ResopnseData;

function UpdatePage(updateView){
    var htmlArr = updateView.HTML;
    if (htmlArr !== undefined){
        let htmlLen = htmlArr.length;
        if (htmlLen != 0){
            for (let i = 0; i < htmlLen; i++){
                let divID = htmlArr[i].Key;
                let html = htmlArr[i].Data;
                document.getElementById(divID).innerHTML = html;
            }
        }
    }

    var scriptArr = updateView.Script;
    if (scriptArr !== undefined){
        let scriptLen = scriptArr.length;
        if (scriptLen != 0){
            const targetelement = "script";
            let allsuspects = document.getElementsByTagName(targetelement);

            for (let i = 0; i < scriptLen; i++){
                let script = scriptArr[i];
                let name = script.Key;
                let data = script.Data;

                //search backwards within nodelist for matching elements to remove
                for (let i=allsuspects.length; i>=0; i--){
                    let target = allsuspects[i];
                    if (target &&
                        target.id ===name){
                        //remove element
                        target.parentNode.removeChild(target);
                        break
                    }
                }

                // add
                let js=document.createElement(targetelement);
                js.setAttribute("type","text/javascript");
                js.id = name;
                js.src = data;
                document.head.appendChild(js);
            }
        }
    }

    var cssArr = updateView.Css;
    if (cssArr !== undefined){
        let cssLen = cssArr.length;
        if (cssLen != 0){
            const targetelement = "link";
            const targetattr = "href";

            let  allsuspects=document.getElementsByTagName(targetelement);

            for (let i = 0; i < cssLen; i++){
                let css = cssArr[i];
                let name = css.Key;
                let data = css.Data;

                for (let i=allsuspects.length; i>=0; i--){ //search backwards within nodelist for matching elements to remove
                    let target = allsuspects[i];
                    if (target &&
                        target.getAttribute(targetattr)!=null &&
                        target.getAttribute(targetattr).indexOf(name)!=-1){
                        //remove element
                        target.parentNode.removeChild(target);
                        break
                    }
                }

                // add
                let style=document.createElement(targetelement);
                style.setAttribute("rel", "stylesheet");
                style.setAttribute("type", "text/css");
                style.id =  name; 
                style.src = data;
                document.head.appendChild(style);
            }
        }
    }

    ResopnseData = updateView.Data;
}