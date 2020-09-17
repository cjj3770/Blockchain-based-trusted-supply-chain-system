<template>
    <div class="app-container">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="80px">
            <el-form-item label="AccountId" prop="userAccountId">
                <el-input v-model="queryParams.userAccountId" placeholder="Please enter UserAccountId" clearable size="small" @keyup.enter.native="handleQuery" />
            </el-form-item>
            
            <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">Search</el-button>
                <!-- <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">Reset</el-button> -->
            </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
            <el-col :span="1.5">
                <el-button
                        type="primary"
                        icon="el-icon-plus"
                        size="mini"
                        @click="handleAdd"
                >Add</el-button>
            </el-col>
            <el-col :span="1.5">
                <el-button
                        type="success"
                        icon="el-icon-edit"
                        size="mini"
                        :disabled="single"
                        @click="handleUpdate"
                >Update</el-button>
            </el-col>
            <el-col :span="1.5">
                <el-button
                        icon="el-icon-refresh"
                        size="mini"
                        @click="handleRefresh"
                >Refresh</el-button>
            </el-col>
        </el-row>

        <el-table v-loading="loading" :data="useraccountList" @selection-change="handleSelectionChange">
            <el-table-column type="index" width="55" align="center" /><el-table-column label="UserAccountId" align="center" prop="accountId" :show-overflow-tooltip="true" />
            <el-table-column label="UserName" align="center" prop="userName" :show-overflow-tooltip="true" />
            <el-table-column label="UserType" align="center" prop="userType" :show-overflow-tooltip="true" />
            <el-table-column label="LastLoginTime" align="center" prop="lastLoginTime" :show-overflow-tooltip="true" />
            <el-table-column label="Status" align="center" prop="status" :show-overflow-tooltip="true" />
            <el-table-column label="Remark" align="center" prop="remark" :show-overflow-tooltip="true" />
            <el-table-column label="Organization ID" align="center" prop="serviceId" :show-overflow-tooltip="true" />
            <el-table-column label="Operation" align="center" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                    <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-edit"
                            @click="handleUpdate(scope.row)"
                    >Update</el-button>
                    <!-- <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-refresh"
                            @click="handleDelete(scope.row)"
                    >Delete</el-button> -->
                </template>
            </el-table-column>
        </el-table>

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
            <el-form ref="form" :model="form" :rules="rules" label-width="120px">
                <el-form-item label="UserAccountId" prop="createUserAccountId"  >
                <el-input v-model="form.createUserAccountId" placeholder="CreateUserAccountId"  />
                </el-form-item>
                
                <el-form-item label="UserName" prop="userName"  >
                <el-input v-model="form.userName" placeholder="UserName"  />
                </el-form-item>
                <el-form-item label="Password" prop="password"  >
                <el-input v-model="form.password" placeholder="Password"  />
                </el-form-item>
                <el-form-item label="UserType" prop="userType"  >
                <el-input v-model="form.userType" placeholder="UserType"  />
                </el-form-item>
                <!-- <el-form-item label="LastLoginTime" prop="lastLoginTime"  >
                <el-input v-model="form.lastLoginTime" placeholder="LastLoginTime"  />
                </el-form-item>
                <el-form-item label="Status" prop="status"  >
                <el-input v-model="form.status" placeholder="Status"  />
                </el-form-item>
                <el-form-item label="Remark" prop="remark"  >
                <el-input v-model="form.remark" placeholder="Remark"  />
                </el-form-item> -->
                <el-form-item label="Organization ID" prop="serviceId"  >
                <el-input v-model="form.serviceId" placeholder="Organization ID"  />
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="submitForm">Confirm</el-button>
                <el-button @click="cancel">Cancel</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex'
    import { listUserAccount, getUserAccount, delUserAccount, addUserAccount, updateUserAccount } from '@/api/useraccount'

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
                useraccountList:[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    userAccountId:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {userAccountId:
                    [
                        {required: true, message: 'UserAccountId不能为空', trigger: 'blur'}
                    ],
                    }
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
                listUserAccount().then(response => {
                        this.useraccountList = response
                        console.log(response)
                        console.log(response.data)
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
                this.form = {
                
                userAccountId: undefined,
                userName: undefined,
                password: undefined,
                userType: undefined,
                // lastLoginTime: undefined,
                // status: undefined,
                // remark: undefined,
                serviceId: undefined,
                }
                this.resetForm('form')
            },
            
            handleRefresh(){
                this.getList()
            },
            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1
                    getUserAccount({"userAccountId": this.queryParams.userAccountId}).then(response => {
                        this.useraccountList = response
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
                this.title = 'Create New UserAccount'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.userAccountId)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleUpdate(row) {
                this.reset()
                const userAccountId = row.userAccountId || this.ids
                getUserAccount(userAccountId).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改userAccount'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function() {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.createUserAccountId !== undefined) {                         
                            addUserAccount({"userAccountId":this.accountId,"createUserAccountId":this.form.createUserAccountId,"userName":this.form.userName,"password":this.form.password,"userType":this.form.userType,"serviceId":this.form.serviceId}).then(response => {                               
                                    this.open = false
                                    
                                    this.getList()
                                    alert("create success")
                            
                            })
                        }
                        
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                const Ids = row.userAccountId || this.ids
                this.$confirm('Are you sure to delete the data item"' + Ids + '"?', 'warning', {
                    confirmButtonText: 'Confirm',
                    cancelButtonText: 'Cancel',
                    type: 'warning'
                }).then(function() {
                    return delUserAccount(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('Delete success')
                }).catch(function() {})
            }
        }
    }
</script>