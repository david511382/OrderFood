var menuData;
var selectedOptionID;
init();

function init(){
    var url = 'api/menu/shopmenu/' + ResopnseData; // from treenode
    $.ajax({
        type:'GET',
        url: url
    }).done(function(data){
        menuData = data;
        selectedOptionID = 0;
        InitShopName();
        initOptionButton();
        initOptionNumSelect();
    });
}

function InitShopName(){
    document.getElementById('shopNameInput').value = menuData.Shop.Name;   
}

function initOptionButton(){
    var optionButtonDiv = document.getElementById('optionButtonDiv');
    var optionTableTr = document.getElementById('optionTableTr');    
    
    if (!menuData.Options){
        return 
    }
    menuData.Options.forEach(function(menuOption) {
        var newTd = document.createElement('td');
        optionTableTr.appendChild(newTd);
        
        // add button
        var newButton = document.createElement('button');
        
        var id= 'none';
        if (menuOption.Option){
            id = menuOption.Option.ID;      
            
            // add remove option button to td
            var newRmButton = document.createElement('button');
            newRmButton.innerHTML = "-";
            newTd.appendChild(newRmButton)
        }
        newButton.Name = id + "OptionButton";
        newButton.innerHTML = menuOption.Name;
        optionButtonDiv.appendChild(newButton);
      }); 

    
    var tdForAddOption = document.createElement('td');
    optionTableTr.appendChild(tdForAddOption);
}

function CreateOptionNumSelect(){
    var optionSelectTd = document.getElementById('optionSelectTd');
    var a = document.createElement('a');
    a.innerHTML = "必選數量";
    optionSelectTd.appendChild(a);


    var select = document.createElement('select');
    select.innerHTML = "必選數量";
    optionSelectTd.appendChild(select);

    var option = document.createElement('option');
    option.value = 0;
    option.innerHTML = "0";
    select.options.add(option);    

    return select;
}

function initOptionNumSelect(){
    if (selectedOptionID == 0){
        var optionSelectTd = document.getElementById('optionSelectTd');
        optionSelectTd.innerHTML = "";
        return;
    }

    var select = CreateOptionNumSelect();
    menuOption = menuData.Options[selectedOptionID];
    menuSelections = menuOption.Selections;
    for (let i = 1; i <= menuSelections.length; i++){
        let option = document.createElement('option');
        option.value = i;
        option.innerHTML = i;
        select.options.add(option);
    }    
}

function newOptionButtonClick(){
    var url = 'manager/newoption';
    $.ajax({
            type:'GET',
            url: url
        }).done(UpdatePage);
}

function newItem(){
    var itemNameInput = document.getElementById('newItemNameInput');
    var name = itemNameInput.value;
    if (!name){
        alert("please input name!");
        return ;
    }

    var itemPriceInput = document.getElementById('newItemPriceInput');
    var price =parseInt(itemPriceInput.value);
    if (isNaN(price)){
        alert("please input integer price");
        return;
    }


    var url = 'api/menu/item';
    var data = {
        shopID: menuData.Shop.ID,
        name: name,
        price: price
    };
    $.ajax({
            type:'POST',
            url: url,
            data: data
        }).done(function(data){
        });
}

// function newSelection(){
//     // @Param optionID formData int true "選單編號"
// // @Param name formData string true "名稱"
// // @Param price formData int false "價格"
// // @Success 200 {object} resp.MenuSelection "菜單"
// // @Failure 500 {string} string "内部错误"
// // @Router /menu/selection [post]

//     var itemNameInput = document.getElementById('newItemNameInput');
//     var name = itemNameInput.value;
//     if (!name){
//         alert("please input name!");
//         return ;
//     }

//     var itemPriceInput = document.getElementById('newItemPriceInput');
//     var price =parseInt(itemPriceInput.value);
//     if (isNaN(price)){
//         alert("please input integer price");
//         return;
//     }

    
//     var url = 'api/menu/item';
//     var data = {
//         shopID: menuData.Shop.ID,
//         name: name,
//         price: price
//     };
//     $.ajax({
//             type:'POST',
//             url: url,
//             data: data
//         }).done(function(data){
//         });
// }

function shopNameInputKeyPress(shopName){
    if (window.event.keyCode==13){
        shopID = menuData.Shop.ID;        
        if (!shopID){
            alert("delete err");
            return false;
        }      

        oldShopName = menuData.Shop.Name;     
        if (!shopName){
            alert("err")
            return false;
        }else if (shopName == ""){
            alert("please input shop name!")
            return false;
        }else if (shopName == oldShopName){
            return false
        }

        UpdateShop(shopID,shopName,
            function(success){
                if (!success){
                    alert('fail');
                    return ;
                }
    
                var url = '/manager/menutree';
                $.ajax({
                    type:'GET',
                    url: url
                }).done(UpdatePage);
    
                alert('修改商店成功!');
            }
        )

        return false;
    };
}

function removeShopButtonClick(){
    shopID = menuData.Shop.ID;        
    if (!shopID){
        alert("delete err");
        return 
    }

    DeleteShop(
        shopID,// from tree node js
        function(success){
            if (!success){
                alert('fail');
                return ;
            }

            var url = '/manager/menutree';
            $.ajax({
                type:'GET',
                url: url
            }).done(UpdatePage);

            toNewShop();
            
            alert('刪除商店成功!');
        }	
    );
}

