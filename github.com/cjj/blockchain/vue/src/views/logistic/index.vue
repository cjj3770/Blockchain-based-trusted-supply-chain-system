<template>
    <div class="app-container">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="80px">
            <el-form-item label="LogisticId" prop="logisticId">
                <el-input v-model="queryParams.logisticId" placeholder="Please enter sLogisticId" clearable size="small" @keyup.enter.native="handleQuery" />
            </el-form-item>
            
            <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">Search</el-button>
                <!-- <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">Empty</el-button> -->
            </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
            <el-col :span="1.5">
                <el-button
                        type="primary"
                        icon="el-icon-plus"
                        size="mini"
                        @click="handleAdd"
                >Create</el-button>
            </el-col>
            <el-col :span="1.5">
                <el-button
                        icon="el-icon-refresh"
                        size="mini"
                        @click="handleQuery"
                >Refresh</el-button>
            </el-col>
        </el-row>

        <el-table v-loading="loading" :data="logisticList" @selection-change="handleSelectionChange">
            <!-- <el-table-column type="selection" width="55" align="center" /> -->
            <el-table-column label="LogisticId" align="center" prop="logisticId" :show-overflow-tooltip="true" />
            <!-- <el-table-column label="CourierAccountId" align="center" prop="courierAccountId" :show-overflow-tooltip="true" /> -->
            <el-table-column label="OrderId" align="center" prop="orderId" :show-overflow-tooltip="true" />
            <!-- <el-table-column label="StoreUserAccountId" align="center" prop="storeUserAccountId" :show-overflow-tooltip="true" />
            <el-table-column label="ManufacturerUserAccountId" align="center" prop="manufacturerUserAccountId" :show-overflow-tooltip="true" /> -->
            
            <el-table-column label="DepartAddress" align="center" prop="departAddress" :show-overflow-tooltip="true" />
            <el-table-column label="DepartTime" align="center" prop="departTime" :show-overflow-tooltip="true" />          
            <el-table-column label="ViaAddress" align="center" prop="viaAddress" :show-overflow-tooltip="true" />
            <el-table-column label="ViaTime" align="center" prop="viaTime" :show-overflow-tooltip="true" />

            <el-table-column label="ArrivalAddress" align="center" prop="arrivalAddress" :show-overflow-tooltip="true" />
            <el-table-column label="ArrivalTime" align="center" prop="arrivalTime" :show-overflow-tooltip="true" />       
            <el-table-column label="LogisticStatus" align="center" prop="logisticStatus" :show-overflow-tooltip="true" />
            <el-table-column label="Operation" align="center" class-name="small-padding" width="250">
                <template slot-scope="scope">
                    <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-edit"
                            @click="handleUpdate(scope.row)"
                            :disabled="scope.row.logisticStatus==='done'"
                    >Update</el-button>
                    <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-success"
                            @click="handleComplete(scope.row)"
                            :disabled="scope.row.logisticStatus==='done'"
                    >Complete</el-button>
                </template>
            </el-table-column>
        </el-table>

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
            <el-form ref="form" :model="form" :rules="rules" label-width="210px">
                
                <el-form-item label="OrderId" prop="orderId"  >
                <el-input v-model="form.orderId" placeholder="OrderId"  />
                </el-form-item>
                <el-form-item label="StoreUserAccountId" prop="storeUserAccountId"  >
                <el-input v-model="form.storeUserAccountId" placeholder="StoreUserAccountId"  />
                </el-form-item>
                <el-form-item label="ManufacturerUserAccountId" prop="manufacturerUserAccountId"  >
                <el-input v-model="form.manufacturerUserAccountId" placeholder="ManufacturerUserAccountId"  />
                </el-form-item>
                <el-form-item label="LogisticId" prop="LogisticId"  >
                <el-input v-model="form.logisticId" placeholder="LogisticId"  />
                </el-form-item>
                <!-- <el-form-item label="ViaAddress" prop="viaAddress"  >
                <el-input v-model="form.viaAddress" placeholder="ViaAddress"  />
                </el-form-item>
                <el-form-item label="ViaTime" prop="viaTime"  >
                <el-input v-model="form.viaTime" placeholder="ViaTime"  />
                </el-form-item> -->
                <!-- <el-form-item label="DepartTime" prop="departTime"  >
                <el-input v-model="form.departTime" placeholder="DepartTime"  />
                </el-form-item>
                <el-form-item label="ArrivalTime" prop="arrivalTime"  >
                <el-input v-model="form.arrivalTime" placeholder="ArrivalTime"  />
                </el-form-item>
                <el-form-item label="DepartAddress" prop="departAddress"  >
                <el-input v-model="form.departAddress" placeholder="DepartAddress"  />
                </el-form-item>
                <el-form-item label="ArrivalAddress" prop="arrivalAddress"  >
                <el-input v-model="form.arrivalAddress" placeholder="ArrivalAddress"  />
                </el-form-item>
                <el-form-item label="LogisticStatus" prop="logisticStatus"  >
                <el-input v-model="form.logisticStatus" placeholder="LogisticStatus"  />
                </el-form-item> -->
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="submitForm">Confirm</el-button>
                <el-button @click="cancel">Cancel</el-button>
            </div>
        </el-dialog>

        <!-- <el-dialog :title="title_update" :visible.sync="open_update" width="500px">
            <el-form ref="form_update" :model="form" :rules="rules" label-width="210px">
            
                <el-form-item label="ViaAddress" prop="viaAddress"  >
                <el-input v-model="form.viaAddress" placeholder="ViaAddress"  />
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="submitForm_update">Confirm</el-button>
                <el-button @click="cancel_update">Cancel</el-button>
            </div>
             
        </el-dialog> -->
    </div>
</template>


<script>
    //import $Scriptjs from '@/api/script'
    //import google from 'google'
    //import {gmapApi} from 'vue2-google-maps'
    import { mapGetters } from 'vuex'
    import { listLogistic, addLogistic,getLogistic, updateLogistic } from '@/api/logistic'
    

    export default {
        name: 'Config',
        data() {
            return {
                // 遮罩层
                loading: true,
                // 选中数组
                ids: [],
                // 非单个禁用
                single: true,
                // 非多个禁用
                multiple: true,
                // 总条数
                total: 0,
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                open_update: false,
                isEdit_update: false,
                
                // 类型数据字典
                typeOptions: [],
                logisticList:[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    logisticId:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {logisticId:
                    [
                        {required: true, message: 'LogisticId不能为空', trigger: 'blur'}
                    ],
                    }
            }
        },
        created() {
            this.getList()
            
        },
        computed: {
            ...mapGetters([
            'accountId',
            'roles',
            'userName',
            'balance'
        ]),
        //google: gmapApi
        },
        methods: {
            /** 查询参数列表 */
            getList() {    
                listLogistic().then(response => {
                        this.logisticList = response
                        this.loading = false
                    }
                )

            },
            // 取消按钮
            cancel() {
                this.open = false
                this.reset()
            },
            // 取消按钮
            cancel_update() {
                this.open_update = false
                this.reset()
            },
            // 表单重置
            reset() {
                this.form = {
                
                orderId: undefined,
                storeUserAccountId: undefined,
                manufacturerUserAccountId: undefined,
                logisticId: undefined,
                courierAccountId: undefined,
                viaAddress: undefined,
                viaTime: undefined,
                departTime: undefined,
                arrivalTime: undefined,
                departAddress: undefined,
                arrivalAddress: undefined,
                logisticStatus: undefined,
                }
                this.resetForm('form')
            },
            

            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1               
                        getLogistic({"logisticId":this.queryParams.logisticId}).then(response => {
                        this.logisticList = response
                    })
            },
            /** 重置按钮操作 */
            resetQuery() {
                this.dateRange = []
                this.resetForm('queryForm')
                this.handleQuery()
            },
            /** 新增按钮操作 */
            handleAdd() {
                //this.reset()
                this.open = true
                this.title = 'Create Logistic'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.logisticId)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },

            
            /** 修改按钮操作 */
            handleUpdate(row) {
                var status;
                var arrivaltime;
                var viatime;
                var viaaddress;
                const courier_id =this.accountId                
                const Ids = row.logisticId
                this.$confirm('Are you sure to update logistic"' + Ids + '"?', 'warning', {
                    confirmButtonText: 'Confirm',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(function() {
                    //get time
                        Date.prototype.Format = function (fmt) { // author: meizz
                        var o = {
                            "M+": this.getMonth() + 1, // 月份
                            "d+": this.getDate(), // 日
                            "h+": this.getHours(), // 小时
                            "m+": this.getMinutes(), // 分
                            "s+": this.getSeconds(), // 秒
                            "q+": Math.floor((this.getMonth() + 3) / 3), // 季度
                            "S": this.getMilliseconds() // 毫秒
                        };
                        if (/(y+)/.test(fmt))
                            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
                        for (var k in o)
                        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
                            return fmt;
                        }
                        var currentDate = new Date().Format("yyyy-MM-dd hh:mm:ss");
                        //get location

                        //Get the latitude and the longitude;
                        function successFunction(position) {
                            var lat = position.coords.latitude;
                            var lng = position.coords.longitude;
                            codeLatLng(lat, lng)
                        }

                        function errorFunction(){
                            alert("Geocoder failed");
                        }

                        function codeLatLng(lat, lng) {
                            var latlng = new google.maps.LatLng({lat: lat, lng: lng});
                            var geocoder = geocoder = new google.maps.Geocoder();
                            geocoder.geocode({ 'latLng': latlng }, function (results, status) {
                        if (status == google.maps.GeocoderStatus.OK) {
                            if (results[1]) {
                            viaaddress=results[1].formatted_address;
                            //alert("Location: " + results[1].formatted_address);
                                                    //check logistic status
                                if(row.logisticStatus=="delivery")
                                    {
                                        status=""
                                        arrivaltime=""
                                        viatime=currentDate

                                    }else{
                                        status="delivery"
                                        arrivaltime=currentDate
                                        viatime=""
                                        viaaddress=""
                                    } 
                       
                                return updateLogistic({"courierAccountId": courier_id,"logisticId":Ids,"viaAddress": viaaddress,"viaTime": viatime,"status":status,"arrivalTime":arrivaltime})
                            }
                            }
                            });
                        }
                        navigator.geolocation.getCurrentPosition(successFunction, errorFunction);
                        // this.loading=true
                        // if (navigator.geolocation) {
                        //     //console.log(navigator.getlocation)
                        //     navigator.geolocation.getCurrentPosition(successFunction, errorFunction);
                        //     this.loading=false
                            
                        // } 

                }).then(() => {
                    this.getList()
                    //alert('update success')
                }).catch(function() {})
            },
            /** 提交按钮 */
            submitForm: function() {
                Date.prototype.Format = function (fmt) { // author: meizz
                        var o = {
                            "M+": this.getMonth() + 1, // 月份
                            "d+": this.getDate(), // 日
                            "h+": this.getHours(), // 小时
                            "m+": this.getMinutes(), // 分
                            "s+": this.getSeconds(), // 秒
                            "q+": Math.floor((this.getMonth() + 3) / 3), // 季度
                            "S": this.getMilliseconds() // 毫秒
                        };
                        if (/(y+)/.test(fmt))
                            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
                        for (var k in o)
                        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
                            return fmt;
                        }
                var currentDate = new Date().Format("yyyy-MM-dd hh:mm:ss"); 
                this.$refs['form'].validate(valid => {
                    if (valid) {
                            addLogistic({"orderId":this.form.orderId,"storeUserAccountId":this.form.storeUserAccountId,"manufacturerUserAccountId":this.form.manufacturerUserAccountId,"logisticId": this.form.logisticId,"courierAccountId": this.accountId,"departTime":currentDate}).then(response => {
                                alert('Create Logistic success')
                                    this.open = false
                                    this.getList()
                            })
                        
                    }
                })
            },
                        
            /** 删除按钮操作 */
            handleComplete(row) {
                const courier_id =this.accountId                
                const Ids = row.logisticId
                this.$confirm('Are you sure to complete logistic"' + Ids + '"?', 'warning', {
                    confirmButtonText: 'Confirm',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(function() {
                    //get time
                        Date.prototype.Format = function (fmt) { // author: meizz
                        var o = {
                            "M+": this.getMonth() + 1, // 月份
                            "d+": this.getDate(), // 日
                            "h+": this.getHours(), // 小时
                            "m+": this.getMinutes(), // 分
                            "s+": this.getSeconds(), // 秒
                            "q+": Math.floor((this.getMonth() + 3) / 3), // 季度
                            "S": this.getMilliseconds() // 毫秒
                        };
                        if (/(y+)/.test(fmt))
                            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
                        for (var k in o)
                        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
                            return fmt;
                        }
                        var currentDate = new Date().Format("yyyy-MM-dd hh:mm:ss");
                    
                    return updateLogistic({"courierAccountId": courier_id,"logisticId":Ids,"viaAddress": "","viaTime": "","status":"done","arrivalTime":currentDate})
        
                }).then(() => {
                    this.getList()
                    //alert('update success')
                    
                }).catch(function() {})
            }
        }
    }
</script>