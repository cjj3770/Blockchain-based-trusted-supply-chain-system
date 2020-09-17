<template>
    <div class="app-container">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
            <el-form-item label="OrderID" prop="orderId">
                <el-input v-model="queryParams.orderId" placeholder="Please enter the orderId" clearable size="small" @keyup.enter.native="handleQuery" />
            </el-form-item> 

            <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">Search</el-button>
                <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">Empty</el-button>
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
                        @click="handleRefresh"
                >Refresh</el-button>
            </el-col>
            <!-- <el-col :span="1.5">
                <el-button
                        type="danger"
                        icon="el-icon-delete"
                        size="mini"
                        :disabled="multiple"
                        @click="handleDelete"
                >Delete</el-button>
            </el-col> -->
        </el-row>

        <el-table v-loading="loading"  :data="orderList" @selection-change="handleSelectionChange">
            <!-- <el-table-column type="index" width="55" align="center" /> -->
            <el-table-column label="orderId" align="center" prop="orderId" :show-overflow-tooltip="true" />
            <!-- <el-table-column label="storeUserAccountId" align="center" prop="storeUserAccountId" :show-overflow-tooltip="true" />
            <el-table-column label="manufacturerUserAccountId" align="center" prop="manufacturerUserAccountId" :show-overflow-tooltip="true" />
            <el-table-column label="logisticUserAccountId" align="center" prop="logisticUserAccountId" :show-overflow-tooltip="true" /> -->
            <el-table-column label="orderName" align="center" prop="orderName" :show-overflow-tooltip="true" />
            <el-table-column label="orderTime" align="center" prop="orderTime" :show-overflow-tooltip="true" />
            <el-table-column label="orderPrice" align="center" prop="orderPrice" :show-overflow-tooltip="true" />
            <!-- <el-table-column label="orderTrace" align="center" prop="orderTrace" :show-overflow-tooltip="true" /> -->
            <el-table-column label="orderStatus" align="center" prop="orderStatus" :show-overflow-tooltip="true" />
            <el-table-column label="departAddress" align="center" prop="departAddress" :show-overflow-tooltip="true" />
            <el-table-column label="arrivalAddress" align="center" prop="arrivalAddress" :show-overflow-tooltip="true" />
            <el-table-column label="Operation" align="center" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                    <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-delete"
                            @click="handleCancel(scope.row)"
                            :disabled="scope.row.orderStatus==='cancelled'||scope.row.orderStatus==='arrived'||scope.row.orderStatus==='expired'||scope.row.orderStatus==='done'"
                    >Cancel</el-button>
                    <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-star-on"
                            @click="handleConfirm(scope.row)"
                            :disabled="scope.row.orderStatus!=='arrived'"
                    >Confirm</el-button>
                </template>
            </el-table-column>
        </el-table>


        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="600px">
            <el-form ref="form" :model="form" :rules="rules" label-width="210px">
                                
                <el-form-item label="orderId" prop="orderId"  >
                <el-input v-model="form.orderId" placeholder="orderId"  />
                </el-form-item>
                <!-- <el-form-item label="storeUserAccountId" prop="storeUserAccountId"  >
                <el-input v-model="form.storeUserAccountId" placeholder="storeUserAccountId"  />
                </el-form-item> -->
                <el-form-item label="manufacturerUserAccountId" prop="manufacturerUserAccountId"  >
                <el-input v-model="form.manufacturerUserAccountId" placeholder="manufacturerUserAccountId"  />
                </el-form-item>
                <el-form-item label="logisticUserAccountId" prop="logisticUserAccountId"  >
                <el-input v-model="form.logisticUserAccountId" placeholder="logisticUserAccountId"  />
                </el-form-item>
                <el-form-item label="orderName" prop="orderName"  >
                <el-input v-model="form.orderName" placeholder="orderName"  />
                </el-form-item>
                <el-form-item label="orderTime" prop="orderTime"  >
                <el-input v-model="form.orderTime" placeholder="orderTime"  />
                </el-form-item>
                <el-form-item label="orderPrice" prop="orderPrice"  >
                <el-input v-model="form.orderPrice" placeholder="orderPrice"  />
                </el-form-item>
                <el-form-item label="traceId" prop="traceId"  >
                <el-input v-model="form.traceId" placeholder="traceId"  />
                </el-form-item>
                <!-- <el-form-item label="orderStatus" prop="orderStatus"  >
                <el-input v-model="form.orderStatus" placeholder="orderStatus"  />
                </el-form-item> -->
                <el-form-item label="departAddress" prop="departAddress"  >
                <el-input v-model="form.departAddress" placeholder="departAddress"  />
                </el-form-item>
                <el-form-item label="arrivalAddress" prop="arrivalAddress"  >
                <el-input v-model="form.arrivalAddress" placeholder="arrivalAddress"  />
                </el-form-item>
                <!-- <el-form-item label="updateStatus" prop="status"  >
                <el-input v-model="form.status" placeholder="updateStatus"  />
                </el-form-item> -->
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="submitForm">Confirm</el-button>
                <el-button @click="cancel">Cancel</el-button>
            </div>
        </el-dialog>



        <!-- <el-dialog :title="title" :visible.sync="open" width="600px">
            <el-form ref="form_cancel" :model="form" :rules="rules" label-width="100px">
                                
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="submitForm">Confirm</el-button>
                <el-button @click="cancel">Cancel</el-button>
            </div>
        </el-dialog> -->
    </div>
</template>

<script>
    import { mapGetters } from 'vuex'
    import { listOrder, getOrder, confirmOrder, addOrder, cancelOrder } from '@/api/order'

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
                // 类型数据字典
                typeOptions: [],
                orderList:[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {}
            }
        },
        computed: {
            ...mapGetters([
            'accountId',
            'roles',
            'userName',
            'balance'
        ])
        },
        created() {
            this.getList()
            
        },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                getOrder({"storeUserAccountId":this.accountId}).then(response => {
                        this.orderList = response
                        console.log(response.orderStatus)
                        this.loading = false
                    }
                )
            },
            // 取消按钮
            cancel() {
                this.open = false
                this.reset()
            },
            // 表单重置
            reset() {
                //this.
            },
            

            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1               
                        getOrder({"storeUserAccountId":this.accountId,"orderId":this.queryParams.orderId}).then(response => {
                        this.orderList = response
                        console.log(response)
                        console.log(response.data)
                    })
            },
                        /** 修改按钮操作 */
            handleConfirm(row) {

                const Ids = row.orderId || this.ids
                this.$confirm('Are you sure to confirm the payment of order"' + Ids + '"?', 'warning', {
                    confirmButtonText: 'Comfirm',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(function() {
                    return confirmOrder({"orderId":Ids,"manufacturerUserAccountId":row.manufacturerUserAccountId,"storeUserAccountId":row.storeUserAccountId,"status":"done"})
                }).then(() => {
                    this.getList()
                    this.msgSuccess('Confirm the payment is successful')
                }).catch(function() {})
                
                
            },
            /** 重置按钮操作 */
            resetQuery() {
                this.queryParams.orderId="";
            },
            /** 新增按钮操作 */
            handleAdd() {
                this.open = true
                this.title = 'Create Order'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.orderId)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleCancel(row) {

                const Ids = row.orderId || this.ids
                this.$confirm('Are you sure to cancel order"' + Ids + '"?', 'warning', {
                    confirmButtonText: 'Comfirm',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(function() {
                    return cancelOrder({"orderId":Ids,"manufacturerUserAccountId":row.manufacturerUserAccountId,"storeUserAccountId":row.storeUserAccountId,"status":"cancelled"})
                }).then(() => {
                    this.getList()
                    this.msgSuccess('Cancel success')
                }).catch(function() {})
                
                
            },
            /** 刷新列表按钮操作 */
            handleRefresh() {
                this.loading = true
                listOrder().then(response => {
                        this.orderList = response
                        this.loading = false
                    }
                )
            },
            /** 提交按钮 */
            submitForm: function() {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.orderId == undefined) {
                            updateOrder(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess('Update success')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addOrder({"orderId":this.form.orderId,"storeUserAccountId":this.accountId,"manufacturerUserAccountId":this.form.manufacturerUserAccountId,"logisticUserAccountId":this.form.logisticUserAccountId,"orderName":this.form.orderName,"orderTime":this.form.orderTime,"orderPrice":parseFloat(this.form.orderPrice),"traceId":this.form.traceId,"departAddress":this.form.departAddress,"arrivalAddress":this.form.arrivalAddress}).then(response => {
                                    //this.msgSuccess('Create success')
                                    alert('Create success')
                                    this.open = false
                                    this.getList()
                                // if (response.code === 200) {
                                //     this.msgSuccess('Create success')
                                //     this.open = false
                                //     this.getList()
                                // } else {
                                //     this.msgError(response.msg)
                                // }
                            })
                        }
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                const Ids = row.orderId || this.ids
                this.$confirm('Are you sure to delete the data item"' + Ids + '"?', 'warning', {
                    confirmButtonText: 'Comfirm',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(function() {
                    return delOrder(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('Delete success')
                }).catch(function() {})
            }
        }
    }
</script>