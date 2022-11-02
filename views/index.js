// var xml = new XMLHttpRequest();

// // login
// var data = JSON.stringify({
//     query: `mutation{
//       login(input: {
//         user_email: "rikianfaisal@gmail.com"
//         user_password: "12345"}) {
//         user{
//           user_id
//           user_name
//           user_image
//           user_status
//           created_date
//           last_update
//           products {
//             UserId
//             ProductId
//             ProductName
//             ProductPrice
//           }
//         }
//         status
//         message
//       }
//     }
//   `
// });

// console.log(window.location.origin)

// xml.open("POST", `http://localhost:8080/query`);

// xml.setRequestHeader("Content-Type", "application/json")

// xml.send(data)

// xml.addEventListener("load", function() {
//     document.querySelector(".user").innerHTML = this.responseText
//     console.log(JSON.parse(this.responseText))
// })

// var formAddProduct = utils.qs("#container-add-product")
            // var containerUpdateProduct = utils.qs("#container-update-product")
            // var containerView = utils.qs("#container-view-product")

            // formAddProduct.style.display = "none"
            // containerUpdateProduct.style.display = "none"
            // containerView.style.display = "none"

            // class Utils {
            //     constructor() {
            //         this.preImgAddProduct = new FileReader()
            //         this.fr = new FileReader()
            //     }

            //     ce(elm) {
            //         return document.createElement(elm)
            //     }

            //     qs(elm) {
            //        return document.querySelector(elm)
            //     }

            //     parseToJSON(jsonString) {
            //         try {
            //             if (typeof jsonString !== "string") return false
            //             return JSON.parse(jsonString)
            //         } catch (error) {
            //             console.error(error.message)
            //             return false
            //         }
            //     }

            //     jsonToString(jsonObj) {
            //         try {
            //             if (typeof jsonObj !== "object") return false
            //             return JSON.stringify(jsonObj)
            //         } catch (error) {
            //             console.error(error.message)
            //             return false
            //         }
            //     }

            //     validationDataPostProduct(formPost) {
            //         if (typeof formPost !== "object") return alert("type form not valid")
            //         // check data product
            //         if (!formPost.get("product_name") || !formPost.get("product_sell") || !formPost.get("product_stock") || !formPost.get("product_info")) {
            //             alert("data product not valid")
            //             return false
            //         }

            //         // check image product
            //         // if (!this.validationImage(formPost.get("product_image"))) {
            //         //     return false
            //         // }

            //         return true
            //     }

            //     validationImage(dataImage) {
            //         if (typeof dataImage !== "object" || !dataImage["size"] || typeof dataImage["size"] !== "number" || !dataImage["type"] || typeof dataImage["type"] !== "string" || !dataImage["type"].match("/")) {
            //             alert("image not valid")
            //             return false
            //         }

            //         if (dataImage["size"] > 100000) {
            //             alert("image too large! receive : " + dataImage["size"])
            //             return false
            //         }

            //         switch (dataImage["type"]) {
            //             case "image/jpg":
            //                 return true
            //             case "image/jpeg":
            //                 return true
            //             case "image/png":
            //                 return true
            //             default:
            //                 alert("type image not valid. receive : " + dataImage["type"])
            //                 return false                        
            //         }
            //     }
            // }

            // class Ajax {
            //     constructor() {
            //         this.xmlDelete = new XMLHttpRequest()
            //         this.xmlUpdate = new XMLHttpRequest()
            //         this.xmlGet = new XMLHttpRequest()
            //         this.xmlPut = new XMLHttpRequest()
            //         this.xmlPost = new XMLHttpRequest()
            //         this.xmlLogout = new XMLHttpRequest()
            //     }

            //     loadData() {
            //         html.openLoading()
            //         console.log("processing load data ...")
            //         var graphQuery = JSON.stringify({
            //             query : "query{user {user{user_id user_email user_name user_image user_status created_date last_update products {UserId ProductId ProductName ProductImage ProductInfo ProductSell ProductStock ProductPrice CreatedDate LastUpdate}}code status}}"
            //         })

            //         this.xmlGet.open("POST", window.origin + "/query", true);
            //         this.xmlGet.setRequestHeader("content-type", "application/json")
            //         return this.xmlGet.send(graphQuery);
            //     }

            //     deleteProduct(id, img) {
            //         // check data product in database
            //         var dataProduct = database.product(id)
            //         if (!dataProduct) {
            //             return alert("data not found")
            //         }
            //         var dataxml = "product_id" + "=" + dataProduct["ProductId"]
            //         this.xmlDelete.open("POST", window.origin + "/products", true)
            //         this.xmlDelete.setRequestHeader("content-type", "application/x-www-form-urlencoded")
            //         this.xmlDelete.send(dataxml)
            //         return
            //     }
                
            //     updateProduct(id) {
            //         // this.productId.innerText = ""
            //         // this.btnResetUpdateProduct.click()
            //         // var product = database.product(id)
            //         // this.productId.innerText = product["ProductId"]
            //         // this.productName.value = product["ProductName"]
            //         // this.productSell.value = product["ProductSell"]
            //         // this.productPrice.value = product["ProductPrice"]
            //         // this.productStock.value = product["ProductStock"]
            //         // this.productInfo.value = product["ProductInfo"]
            //         // this.productImage.innerText = product["ProductImage"]
            //     }

            //     logout() {
            //         this.xmlLogout.open("PUT", window.origin + "/user/logout/", true)
            //         return this.xmlLogout.send()
            //     }
            // }

            // class Database{
            //     constructor() {
            //         this.products = []
            //     }

            //     product(id) {
            //         var product = this.products.find(function(data) {
            //             return data["ProductId"] === id
            //         })

            //         return product
            //     }

            //     updateProduct(dataProduct) {
            //         for (var i = 0; i < this.products.length; i++) {
            //             if (this.products[i]["ProductId"] === dataProduct["ProductId"]) {
            //                 this.products[i]["ProductInfo"] = dataProduct["ProductInfo"]
            //                 this.products[i]["ProductName"] = dataProduct["ProductName"]
            //                 this.products[i]["ProductPrice"] = dataProduct["ProductPrice"]
            //                 this.products[i]["ProductSell"] = dataProduct["ProductSell"]
            //                 this.products[i]["ProductStock"] = dataProduct["ProductStock"]
            //                 this.products[i]["CreatedDate"] = dataProduct["CreatedDate"]
            //                 this.products[i]["LastUpdate"] = dataProduct["LastUpdate"]
            //                 break
            //             }
                        
            //         }
            //     }

            //     insertProduct(dataProduct) {
            //         if (typeof dataProduct !== "object") {
            //             alert("failed insert data product to master")
            //             return false
            //         }

            //         this.products.push(dataProduct)
            //         return true
            //     }
            // }

            // class Profile {
            //     constructor() {
            //         this.userName = utils.qs("#user-name")
            //     }
            // }

            // class Html {
            //     constructor() {
            //         // loading
            //         this.loading = utils.qs(".loading")
            //         // img preview add product
            //         this.conImgPrevAddProduct = utils.qs("#con-img-prev-add-product")
            //         this.imgAddProduct = utils.qs("#product_image")

            //         this.buttonMasterData = utils.qs("#button-master-data")

            //         // container master data
            //         this.containerMasterData = utils.qs("#container-master-data")

            //         // master body product
            //         this.masterBodyProduct = utils.qs("#body-master-product")

            //         // form post new product
            //         this.formPostNewDataProduct = utils.qs("#formDataProduct")
            //         this.btnSendPostNewDataProduct = utils.qs("#btn_add_new_product")

            //         // form update product
            //         this.formUpdateDataProduct = utils.qs("#formUpdateDataProduct")
            //         this.productImagePreview = utils.qs("#up_prev_product_image")
            //         this.productId = utils.qs("#up_product_id")
            //         this.productName = utils.qs("#up_product_name")
            //         this.productSell = utils.qs("#up_product_sell")
            //         this.productPrice = utils.qs("#up_product_price")
            //         this.productStock = utils.qs("#up_product_stock")
            //         this.productInfo = utils.qs("#up_product_info")
            //         this.productImage = utils.qs("#up_product_image")

            //         this.btnResetFormUpdateProduct = utils.qs("#btn_up_reset")
            //         this.btnSendUpdateProduct = utils.qs("#btn_up_product")
            //         this.btnCloseUpdate = utils.qs("#btn_up_close")

            //         // action add product
            //         this.buttonAddProduct = utils.qs("#add-product")
            //         this.formAddProduct = utils.qs("#container-add-product")
            //         this.btnResetAddNewProduct = utils.qs("#btn_reset_add_new_product")
            //         this.btnCloseAddNewProduct = utils.qs("#btn_close_add_new_product")

            //         // action update product
            //         this.containerUpdateProduct = utils.qs("#container-update-product")

            //         // action for view listener
            //         this.containerView = utils.qs("#container-view-product")
            //         this._productName = utils.qs("#view-name-product")
            //         this._productInfo = utils.qs("#view-info-product")
            //         this._productImage = utils.qs("#view-image-product")
            //         this._productPrice = utils.qs("#view-price-product")
            //         this._productSell = utils.qs("#view-sell-product")
            //         this._viewBactToMaster = utils.qs("#view-bact-to-master")

            //         // logout
            //         this.btnLogout1 = utils.qs("#btn-logout1")
            //         this.btnLogout2 = utils.qs("#btn-logout2")
            //     }

            //     resetDataAddProduct() {
            //         this.btnResetAddNewProduct.click()
            //         this.conImgPrevAddProduct.innerHTML = ""
            //     }

            //     createContainerProduct(dataProduct) {
            //         "use strict"
            //         var containerProduct = utils.ce("tr")
            //         containerProduct.setAttribute("id", "up" + dataProduct["ProductId"])
                    
            //         var productId = utils.ce("td")
            //         productId.innerText = dataProduct["ProductId"]
            //         containerProduct.appendChild(productId)

            //         var productName = utils.ce("td")
            //         productName.innerText = dataProduct["ProductName"]
            //         containerProduct.appendChild(productName)
                    
            //         // var productImage = utils.ce("td")
            //         // containerProduct.appendChild(productImage)

            //         // var imgProduct = utils.ce("img")
            //         // imgProduct.src = "http://192.168.43.38:9092/cdn-lawnsoor/media/image-product/" + dataProduct["ProductImage"]
            //         // imgProduct.setAttribute("loading", "lazy")
            //         // productImage.appendChild(imgProduct)
                    
            //         var productPrice = utils.ce("td")
            //         productPrice.innerText = dataProduct["ProductPrice"]
            //         containerProduct.appendChild(productPrice)
                    
            //         var productSell = utils.ce("td")
            //         productSell.innerText = dataProduct["ProductSell"]
            //         containerProduct.appendChild(productSell)
                    
            //         var productStock = utils.ce("td")
            //         productStock.innerText = dataProduct["ProductStock"]
            //         containerProduct.appendChild(productStock)

            //         // var productInfo = utils.ce("td")
            //         // productInfo.innerText = dataProduct["ProductInfo"]
            //         // containerProduct.appendChild(productInfo)
                    
            //         // var createdDate = utils.ce("td")
            //         // createdDate.innerText = dataProduct["CreatedDate"]
            //         // containerProduct.appendChild(createdDate)
                    
            //         // var lastUpdate = utils.ce("td")
            //         // lastUpdate.innerText = dataProduct["LastUpdate"]
            //         // containerProduct.appendChild(lastUpdate)
                    
            //         var actions = utils.ce("td")
            //         actions.setAttribute("class", "con-actions")
            //         containerProduct.appendChild(actions)

            //         var buttonViewAction = utils.ce("button")
            //         buttonViewAction.setAttribute("class", "btn btn-success")
            //         buttonViewAction.innerText = "View"
            //         buttonViewAction.addEventListener("click", function() {
            //             return html.openContainerViewProduct(dataProduct)
            //         })
            //         actions.appendChild(buttonViewAction)
                    
            //         var buttonUpdateAction = utils.ce("button")
            //         buttonUpdateAction.setAttribute("class", "btn btn-warning")
            //         buttonUpdateAction.innerText = "Update"
            //         buttonUpdateAction.addEventListener("click", function() {
            //             html.openContainerUpdateProduct()
            //             return html.openFormUpdateProduct(dataProduct["ProductId"])
            //         })

            //         actions.appendChild(buttonUpdateAction)

            //         var buttonDeleteAction = utils.ce("button")
            //         buttonDeleteAction.setAttribute("class", "btn btn-danger")
            //         buttonDeleteAction.innerText = "Delete"
            //         actions.appendChild(buttonDeleteAction)

            //         buttonDeleteAction.addEventListener("click", function() {
            //             var confirm = prompt("Perhatian! Data yang dihapus tidak dapat dikembalikan!\n\nketik 'yes' untuk melanjutkan")
            //             if (confirm !== "yes") return
            //             return ajax.deleteProduct(dataProduct["ProductId"], dataProduct["ProductImage"])
            //         })
                    
            //         return containerProduct
            //     }

            //     openFormUpdateProduct(id) {
            //         "use strict"
            //         this.productId.innerText = ""
            //         this.productImage.innerText = ""
            //         this.btnResetFormUpdateProduct.click()
                    
            //         var dataProduct = database.product(id)
                    
            //         if (!dataProduct) return alert("data product not found")
            //         this.productImagePreview.src = "http://192.168.43.38:9095/cdn-lawnsoor/media/image-product/" + dataProduct["ProductImage"]
            //         this.productId.innerText = dataProduct["ProductId"]
            //         this.productName.value = dataProduct["ProductName"]
            //         this.productSell.value = dataProduct["ProductSell"]
            //         this.productPrice.value = dataProduct["ProductPrice"]
            //         this.productStock.value = dataProduct["ProductStock"]
            //         this.productInfo.value = dataProduct["ProductInfo"]
            //         this.productImage.innerText = dataProduct["ProductImage"]
            //         return
            //     }

            //     openFormAddProduct() {
            //         this.containerMasterData.style.display = "none"
            //         this.formAddProduct.style.display = "block"
            //         this.containerUpdateProduct.style.display = "none"
            //         this.containerView.style.display = "none"
            //         return
            //     }
                
            //     openMasterData() {
            //         this.containerMasterData.style.display = "block"
            //         this.formAddProduct.style.display = "none"
            //         this.containerUpdateProduct.style.display = "none"
            //         this.containerView.style.display = "none"
            //         return
            //     }

            //     openContainerUpdateProduct() {
            //         this.containerUpdateProduct.style.display = "block"
            //         this.formAddProduct.style.display = "none"
            //         this.containerMasterData.style.display = "none"
            //         this.containerView.style.display = "none"
            //     }

            //     openContainerViewProduct(dataProduct) {
            //         this.containerView.style.display = "flex"
            //         this.containerUpdateProduct.style.display = "none"
            //         this.formAddProduct.style.display = "none"
            //         this.containerMasterData.style.display = "none"
            //         return this.passingDataToContainerProduct(dataProduct)
            //     }

            //     passingDataToContainerProduct(dataProduct, userId) {
            //         this._productName.innerText = dataProduct["ProductName"]
            //         this._productInfo.innerText = dataProduct["ProductInfo"]
            //         this._productPrice.innerText = "Rp. " + dataProduct["ProductPrice"].toLocaleString("id-ID") + ",-"
            //         this._productSell.innerText = "Rp. " + dataProduct["ProductSell"].toLocaleString("id-ID") + ",-"
            //         var imgParts = dataProduct["ProductImage"].split("/")
            //         this._productImage.src = "http://192.168.43.38:9095/cdn-lawnsoor/media/image-product/" + dataProduct["ProductImage"]
            //     }
                
            //     clearContainerView() {
            //         this._productName.innerText = ""
            //         this._productInfo.innerText = ""
            //         this._productImage.src = ""
            //         this._productPrice.innerText = ""
            //         this._productSell.innerText = ""
            //     }

            //     openLoading() {
            //         this.loading.style.display = "block"
            //     }

            //     closeLoading() {
            //         this.loading.style.display = "none"
            //     }
            // }
            
            // var ajax = new Ajax()
            // var utils = new Utils()
            // var database = new Database()
            // var profile = new Profile()
            // var html = new Html()

            // html.formAddProduct.style.display = "none"
            // html.containerUpdateProduct.style.display = "none"
            // html.containerView.style.display = "none"

            // html._viewBactToMaster.addEventListener("click", function() {
            //     // clear container view
            //     html.clearContainerView()
            //     html.openMasterData()
            // })

            // // action container add product
            // html.btnCloseAddNewProduct.addEventListener("click", function(e) {
            //     e.preventDefault()
            //     html.openMasterData()
            // })

            // html.imgAddProduct.addEventListener("change", function() {
            //     if (!utils.validationImage(this.files[0])) {
            //         return
            //     }

            //     html.conImgPrevAddProduct.innerHTML = ""

            //     utils.preImgAddProduct.readAsDataURL(this.files[0])
            // })

            // utils.preImgAddProduct.addEventListener("load", function() {
            //     var imgPrev = utils.ce("img")
            //     imgPrev.style.width = "100%"
            //     imgPrev.style.marginBottom = "10px"
            //     imgPrev.style.borderRadius = "5px"
            //     imgPrev.src = this.result
            //     html.conImgPrevAddProduct.appendChild(imgPrev)
            // })

            // html.btnResetAddNewProduct.addEventListener("click", function() {
            //     html.conImgPrevAddProduct.innerHTML = ""
            // })

            // html.buttonAddProduct.addEventListener("click", function() {
            //     html.openFormAddProduct()
            // })

            // // open master data
            // html.buttonMasterData.addEventListener("click", function() {
            //     html.openMasterData()
            // })

            // // get all data
            // ajax.loadData()
            // ajax.xmlGet.addEventListener("load", function() {
            //     var data = utils.parseToJSON(this.responseText)
                
            //     if (!data) {
            //         html.closeLoading()
            //         return alert("data not found")
            //     }
                
            //     profile.userName.innerText = data["data"]["user"]["user"]["user_name"]
                
            //     var dataProducts = data["data"]["user"]["user"]["products"]

            //     if (dataProducts === null || !dataProducts) {
            //         html.closeLoading()
            //         return
            //     }
                
            //     database.products = dataProducts

            //     for (let i = 0; i < dataProducts.length; i++) {
            //         html.masterBodyProduct.appendChild(html.createContainerProduct(dataProducts[i]))
            //     }

            //     html.closeLoading()
            //     return console.log("load data finish...")
            // })

            // // update event
            // html.btnCloseUpdate.addEventListener("click", function(e) {
            //     e.preventDefault()
            //     html.openMasterData()
            // })
            
            // html.btnSendUpdateProduct.addEventListener("click", function(e) {
            //     e.preventDefault()
            //     // check data product exist or not in database
            //     var product = database.product(html.productId.innerText)
            //     if (!product) return alert("product not valid")
            //     var dataProduct = new FormData(html.formUpdateDataProduct)
            //     dataProduct.set("product_id", product["ProductId"])
            //     dataProduct.set("product_image", product["ProductImage"])
            //     dataProduct.set("created_date", product["CreatedDate"])

            //     ajax.xmlUpdate.open("PUT", window.location.origin + "/products/", true)
            //     ajax.xmlUpdate.send(dataProduct)
            //     console.log("processing update data product...")
            //     return
            // })

            // ajax.xmlUpdate.addEventListener("load", function(e) {
            //     var responseUpdateProduct = utils.parseToJSON(this.responseText)
            //     if (!responseUpdateProduct) return alert(this.responseText)
            //     var conProduct = utils.qs("#up" + responseUpdateProduct["product"]["ProductId"])
            //     if (!conProduct || conProduct === null) {
            //         alert("container product not found")
            //         return window.location.reload()
            //     }
            //     conProduct.remove()
            //     database.updateProduct(responseUpdateProduct["product"])
            //     html.masterBodyProduct.appendChild(html.createContainerProduct(responseUpdateProduct["product"]))
            //     alert("success update data product...")
            //     html.openMasterData()
            //     html.btnResetFormUpdateProduct.click()
            //     var cp = utils.qs("#up"+ responseUpdateProduct["product"]["ProductId"])
            //     cp.style.backgroundColor = "green"
            //     var borderTime = setTimeout(function() {
            //         cp.style.backgroundColor = ""
            //         return clearTimeout(borderTime)
            //     }, 3000)

            //     window.location = window.location.origin + "/#up" + responseUpdateProduct["product"]["ProductId"]
            //     return
            // })

            // // post new data product
            // var formNewProduct = null
            // html.btnSendPostNewDataProduct.addEventListener("click", function(e) {
            //     e.preventDefault()
            //     html.openLoading()
            //     if (formNewProduct !== null) formNewProduct = null
            //     formNewProduct = new FormData(html.formPostNewDataProduct)
            //     // formNewProduct.set()
            //     if (!utils.validationDataPostProduct(formNewProduct)) {
            //         return html.closeLoading()
            //     }

            //     if (!html.imgAddProduct.files[0]) {
            //         alert("please chooce image product")
            //         return html.closeLoading()
            //     }

            //     return utils.fr.readAsArrayBuffer(html.imgAddProduct.files[0])
            // })
            
            // utils.fr.addEventListener("load", function(e) {
            //     var target = window.location.protocol + "//" + window.location.hostname + ":9095" + "/products/images/"
            //     ajax.xmlPost.open("POST", target, true)
            //     ajax.xmlPost.withCredentials = true
            //     ajax.xmlPost.setRequestHeader("Content-Type", html.imgAddProduct.files[0].type)
            //     ajax.xmlPost.setRequestHeader("token", "image/jpeg")
            //     ajax.xmlPost.send(this.result)
            //     return
            // })

            // ajax.xmlPost.addEventListener("load", function(e) {
            //     var postMessage = utils.parseToJSON(this.responseText)
            //     if (!postMessage) return html.closeLoading()
            //     if (postMessage["method"] === "upload image") {
            //         if (postMessage["status"] !== "ok") {
            //             alert(postMessage["error"])
            //             return html.closeLoading()
            //         }
            //         console.log(utils.parseToJSON(this.responseText))
            //         formNewProduct.set("product_id", postMessage["productID"])
            //         formNewProduct.set("enctype", "application/x-www-form-urlencoded")
            //         formNewProduct.forEach(v => console.log(v))
            //         ajax.xmlPost.open("POST", window.location.origin + "/user/products", true)
            //         ajax.xmlPost.send(formNewProduct)
            //         return html.closeLoading()
            //     }

            //     if (postMessage["method"] === "add new product" && postMessage["status"] === "ok") {
            //         console.log(postMessage)
            //         if(!database.insertProduct(postMessage["info"])) {
            //             html.closeLoading()
            //             window.location.reload()
            //             return 
            //         }
            //         html.masterBodyProduct.appendChild(html.createContainerProduct(postMessage["info"]))

            //         html.closeLoading()
            //         alert("success post product...")
                    
            //         html.openMasterData()
            //         html.resetDataAddProduct()

            //         var cp = utils.qs("#up"+ postMessage["info"]["ProductId"])
            //         cp.style.backgroundColor = "green"
            //         var borderTime = setTimeout(function() {
            //             cp.style.backgroundColor = ""
            //             return clearTimeout(borderTime)
            //         }, 3000)

            //         window.location = window.location.origin + "/#up" + postMessage["info"]["ProductId"]
            //         return
            //     }

            //     console.log(postMessage)
            //     return html.closeLoading()
            // })

            // // delete data product
            // ajax.xmlDelete.addEventListener("load", function() {
            //     var response = utils.parseToJSON(this.responseText)
            //     console.log(response)
            //     if (!response) return alert(this.responseText)
            //     var product = utils.qs("#up" + response["ProductId"])
            //     if (!product) return alert("product not found")
            //     return product.remove()
            // })
        
            // // logout
            // html.btnLogout1.addEventListener("click", function () {
            //     console.log("process log out")
            // })

            // html.btnLogout2.addEventListener("click", function () {
            //     ajax.logout()
            // })

            // ajax.xmlLogout.addEventListener("load", function() {
            //     var message = utils.parseToJSON(this.responseText)
            //     if (!message) {
            //         return
            //     }

            //     console.log(message)
            //     window.location.reload()
            // })