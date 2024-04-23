function initProductTable(keyword)
{
    productTable = $('#productTable').DataTable({
        language: {
            url: "/static/vendor/dataTables/languages/ch-tw.json"  
        },
        oLanguage: {
            "sProcessing": "<i class='fa fa-spinner fa-spin fa-3x fa-fw'></i><span class='sr-only'>Loading...</span>"
        },
        ordering: false,
        processing: true,
        serverSide:true,
        ajax: {
            url: "/api/products/getList",
            type: "GET",
        },
        search: {
            "search": keyword
        },
        columns:[
            {
                data: 'name', 
            },
            {
                data: null,
                render: function(data, type, row) {
                    output = '<div style="display:flex; align-items:center;">';
                    // - button
                    output += '<button type="button" class="btn" style="padding: 5px 10px;" onclick="amountHandler(\'sub\', '+row.id+')"><i class="fa-solid fa-minus"></i></button>';
                    // input
                    output += '<input type="number" class="amountInput form-control" style="width:100px;" data-id="'+row.id+'" value="'+row.amount+'">';
                    // + button
                    output += '<button type="button" class="btn" style="padding: 5px 10px;" onclick="amountHandler(\'add\', '+row.id+')"><i class="fa-sharp fa-solid fa-plus"></i></button>';
                    // updateBtn
                    output += '<button type="button" class="btn btn-success btn-sm mdfBtn_'+row.id+'" style="margin-left:5px" onclick="updateAmount(this, '+row.id+')">修改</button>';
                    // loading icon
                    output += '<i class="fa fa-spinner fa-spin loading_'+row.id+'" style="margin-left:5px;font-size:20px;display:none;"></i>';
                    // success
                    output += '<i class="fa fa-solid fa-check statusIcon success_'+row.id+'" style="color:#12af55;margin-left:5px;font-size:20px;display:none;"></i>';
                    // failed
                    output += '<i class="fa fa-solid fa-xmark statusIcon failed_'+row.id+'" style="color:#FF0000;margin-left:5px;font-size:20px;display:none;"></i>';
                    output += '</div>';
                    output += '<div class="msgItem msgItem_'+row.id+'" style="display:none;">message: <span class="msg_'+row.id+'" style="color:#FF0000;"></span></div>';
                    return output;
                }
            },
            {
                data: 'amountNotice', 
                render: function(data, type, row) {
                    if(data>0) {
                        if(row.amount <= row.amountNotice) {
                            return '<span style="color:#FF0000;">' + data + '</span>';
                        } else {
                            return data;
                        }
                    } else {
                        return "無";
                    }
                }
            },
            {
                data: 'price',
            },
            {
                data: 'discount_price',
                render: function(data, type, row) {
                    if(parseInt(data) > 0) {
                        return data;
                    }
                    return "-";
                }
            },
            {
                data: 'tname',
                render: function(data, type, row) {
                    if(row.tname.length > 0) {
                        return row.tname;
                    } else {
                        return "無";
                    }
                }
            },
            {
                data: 'formatExpiredDate',
                render: function(data, type, row) {
                    if(!data) {
                        return "-";
                    } else {
                        // 小於三個月顯示紅字
                        let today = new Date();
                        let threeMonthsAgo = new Date(
                            today.getFullYear(), 
                            today.getMonth() + 3, 
                            today.getDate()
                        );
                        if(new Date(data).getTime() < threeMonthsAgo.getTime()) {
                            return "<span style='color: red'>" + data + "</span>";
                        }
                        return data;
                    }
                }
            },
            {
                data: 'formatTime',
            },
            {
                data: null, 
                render: function(data, type, row) {
                    hasPicture = '';
                    if(row.pictureCnt > 0) {
                        hasPicture = 'hasPicture';
                    }
                    output = `<div style="display:flex; align-items:center;height:30px;">
                                <i title="圖片管理" class="fa-sharp fa-solid fa-images `+hasPicture+`" style="cursor:pointer;font-size:20px;margin-right:10px;" onclick="imagesManage(`+row.id+`)"></i>
                                <i class="fa-solid fa-pen-to-square" style="cursor:pointer;font-size:20px;" onclick="modifyProduct(`+row.id+`)"></i>
                                <i class="fa-sharp fa-solid fa-trash" style="cursor:pointer;font-size:20px;margin-left:10px;" onclick="doDelete(this, `+row.id+`)"></i>
                            </div>`;
                    return output;
                }
            },
        ],
        columnDefs:[
            {
                targets: 6,
                orderable: false,
                responsivePriority: 2,
            },
            {
                targets: 0,
                responsivePriority: 1,
            }
        ],
    });
    return productTable;
}

function initOrdersTable()
{
    orderTable = $('#orderTable').DataTable({
        language: {
            url: "/static/vendor/dataTables/languages/ch-tw.json"  
        },
        oLanguage: {
            "sProcessing": "<i class='fa fa-spinner fa-spin fa-3x fa-fw'></i><span class='sr-only'>Loading...</span>"
        },
        processing: true,
        serverSide:true,
        order: [
            [6, 'desc']
        ],
        ajax: {
            url: "/api/orders/getLists",
            type: "GET",
        },
        columns:[
            {
                data: 'name',
            },
            {
                data: 'contact', 
            },
            {
                data: 'total', 
            },
            {
                data: null,
                className: 'none',
                render : function(data, type, row) {
                    output = '';
                    if(row.detail !== null) {
                        output += `<table class="table table-bordered">
                        <thead>
                            <tr>
                                <th>名稱</th>
                                <th>數量</th>
                                <th>單價</th>
                                <th>總價</th>
                            </tr>
                        </thead>
                        <tbody>`;
                        for (const index in row.detail) {
                            output += `<tr>
                                <td>` + row.detail[index].pname + `</td>
                                <td>` + row.detail[index].amount + `</td>
                                <td>` + row.detail[index].price + `</td>
                                <td>` + row.detail[index].total + `</td>
                            </tr>`;
                        }
                        output += `</tbody></table>`;
                    } else {
                        output += '無';
                    }
                    return output;
                }
            },
            {
                data: 'address', 
            },
            {
                data: 'remark', 
            },
            {
                data: 'formatTime', 
            },
            {
                data: null,
                render: function(data, type, row) {
                    output = '<div style="display:flex; align-items:center;">';
                    output += '<i class="fa-solid fa-pen-to-square" style="cursor:pointer;font-size:20px;" onclick="modifyOrder('+row.id+')"></i>';
                    output += '<i class="fa-sharp fa-solid fa-trash" style="cursor:pointer;font-size:20px;margin-left:10px;" onclick="doDelete(this, '+row.id+')"></i>';
                    output += '</div>';
                    return output;
                }
            }
        ],
        columnDefs:[
            {
                targets: 2,
                orderable: false,
            },
            {
                targets: 7,
                orderable: false,
                responsivePriority: 2,
            },
        ],
    });
    return orderTable;
}

function initPorductsPicTable(id)
{
    productsPicTable = $('#productsPicTable').DataTable({
        searching: false,
        ordering: false,
        language: {
            url: "/static/vendor/dataTables/languages/ch-tw.json"  
        },
        oLanguage: {
            "sProcessing": "<i class='fa fa-spinner fa-spin fa-3x fa-fw'></i><span class='sr-only'>Loading...</span>"
        },
        processing: true,
        serverSide:true,
        ajax: {
            url: "/api/products/getPictures?pid=" + id,
            type: "GET",
        },
        columns:[
            {
                data: 'id',
            },
            {
                data: 'picture', 
                render: function(data, type, row) {
                    output = `
                    <img src="/static/uploads/`+row.pid+`/`+row.picture+`" style="width:150px;">`;
                    return output;
                }
            },
            {
                data: 'formatTime', 
            },
            {
                data: null,
                render: function(data, type, row) {
                    output = `<div style="display:flex; align-items:center;height:30px;">
                                <i class="fa-solid fa-pen-to-square" style="cursor:pointer;font-size:20px;" onclick="modify(`+row.id+`, `+row.pid+`, '`+row.picture+`')"></i>
                                <i class="fa-sharp fa-solid fa-trash" style="cursor:pointer;font-size:20px;margin-left:10px;" onclick="doDelete(`+row.id+`)"></i>
                            </div>`;
                    return output;
                }
            }
        ],
        columnDefs:[
            {
                targets: 1,
                responsivePriority: 1,
            },
        ],
    });
    return productsPicTable;
}


function initPorductsTypeTable()
{
    productTypeTable = $('#productTypeTable').DataTable({
        searching: false,
        ordering: false,
        language: {
            url: "/static/vendor/dataTables/languages/ch-tw.json"  
        },
        oLanguage: {
            "sProcessing": "<i class='fa fa-spinner fa-spin fa-3x fa-fw'></i><span class='sr-only'>Loading...</span>"
        },
        processing: true,
        serverSide:true,
        ajax: {
            url: "/api/products/type",
            type: "GET",
        },
        columns:[
            {
                data: 'id',
            },
            {
                data: 'name', 
            },
            {
                data: 'formatCreateTime', 
            },
            {
                data: 'formatUpdateTime', 
            },
            {
                data: null,
                render: function(data, type, row) {
                    output = `<div style="display:flex; align-items:center;height:30px;">
                                <i class="fa-solid fa-pen-to-square" style="cursor:pointer;font-size:20px;" onclick="modify(`+row.id+`, '`+row.name+`')"></i>
                                <i class="fa-sharp fa-solid fa-trash" style="cursor:pointer;font-size:20px;margin-left:10px;" onclick="doDelete(`+row.id+`)"></i>
                            </div>`;
                    return output;
                }
            }
        ],
        columnDefs:[
            {
                targets: 1,
                responsivePriority: 1,
            },
        ],
    });
    return productTypeTable;
}

function initBannerTable()
{
    bannerTable = $('#bannerTable').DataTable({
        searching: false,
        ordering: false,
        language: {
            url: "/static/vendor/dataTables/languages/ch-tw.json"  
        },
        oLanguage: {
            "sProcessing": "<i class='fa fa-spinner fa-spin fa-3x fa-fw'></i><span class='sr-only'>Loading...</span>"
        },
        processing: true,
        serverSide:true,
        ajax: {
            url: "/api/banner",
            type: "GET",
        },
        columns:[
            {
                data: 'picture', 
                render: function(data, type, row) {
                    output = `
                    <img src="/static/uploads/banner/`+row.picture+`" style="width:150px;">`;
                    return output;
                }
            },
            {
                data: 'formatTime', 
            },
            {
                data: null,
                render: function(data, type, row) {
                    output = `<div style="display:flex; align-items:center;height:30px;">
                                <i class="fa-solid fa-pen-to-square" style="cursor:pointer;font-size:20px;" onclick="modify(`+row.id+`,'`+row.picture+`')"></i>
                                <i class="fa-sharp fa-solid fa-trash" style="cursor:pointer;font-size:20px;margin-left:10px;" onclick="doDelete(`+row.id+`)"></i>
                            </div>`;
                    return output;
                }
            }
        ],
        columnDefs:[
            {
                targets: 1,
                responsivePriority: 1,
            },
        ],
    });
    return bannerTable;
}


function initContactTable()
{
    contactTable = $('#contactTable').DataTable({
        searching: false,
        ordering: false,
        language: {
            url: "/static/vendor/dataTables/languages/ch-tw.json"  
        },
        oLanguage: {
            "sProcessing": "<i class='fa fa-spinner fa-spin fa-3x fa-fw'></i><span class='sr-only'>Loading...</span>"
        },
        processing: true,
        serverSide:true,
        ajax: {
            url: "/api/contact",
            type: "GET",
        },
        columns:[
            {
                data: 'name', 
            },
            {
                data: 'contact', 
            },
            {
                data: 'message', 
            },
            {
                data: 'formatTime', 
            },
            {
                data: null,
                render: function(data, type, row) {
                    output = `<div style="display:flex; align-items:center;height:30px;">
                                <i class="fa-sharp fa-solid fa-trash" style="cursor:pointer;font-size:20px;margin-left:10px;" onclick="doDelete(`+row.id+`)"></i>
                            </div>`;
                    return output;
                }
            }
        ],
        columnDefs:[
            {
                targets: 1,
                responsivePriority: 1,
            },
        ],
    });
    return contactTable;
}
