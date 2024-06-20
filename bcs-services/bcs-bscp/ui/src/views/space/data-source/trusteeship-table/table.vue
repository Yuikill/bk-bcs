<template>
  <bk-loading style="min-height: 300px" :loading="tableLoading">
    <bk-table
      :data="tableData"
      :border="['outer']"
      :remote-pagination="true"
      :pagination="pagination"
      @page-limit-change="handlePageLimitChange"
      @page-value-change="handlePageCurrentChange">
      <bk-table-column :label="$t('表格名称')">
        <template #default="{ row }">
          <bk-button v-if="row.name" text theme="primary"> {{ row.name }} </bk-button>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('表格描述')" prop="memo" />
      <bk-table-column :label="$t('关联配置项')">
        <template #default="{ row }">
          <bk-button v-if="row.name" text theme="primary" :disabled="row.config === 0"> {{ row.config }} </bk-button>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('最近更新人')" prop="updator" />
      <bk-table-column :label="$t('最近更新时间')" prop="updatedAt" />
      <bk-table-column :label="$t('操作')">
        <template #default="{ row }">
          <div class="action-btns">
            <bk-button text theme="primary" @click="handleEditData(row)">{{ $t('编辑数据') }}</bk-button>
            <bk-button text theme="primary" @click="handleDeleteDataSource(row)">{{ $t('编辑表结构') }}</bk-button>
            <bk-popover
              theme="light trusteeship-table-actions-popover"
              placement="bottom-end"
              :popover-delay="[0, 100]"
              :arrow="false">
              <div class="more-actions">
                <Ellipsis class="ellipsis-icon" />
              </div>
              <template #content>
                <div class="config-actions">
                  <div class="action-item" @click="handleImportTable(row)">{{ $t('导入表格') }}</div>
                  <div class="action-item" @click="handleExportTable(row)">{{ $t('导出表格') }}</div>
                  <div :class="['action-item', { disabled: row.config !== 0 }]" @click="handleDeleteTable(row)">
                    {{ $t('删除') }}
                  </div>
                </div>
              </template>
            </bk-popover>
          </div>
        </template>
      </bk-table-column>
    </bk-table>
  </bk-loading>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Ellipsis } from 'bkui-vue/lib/icon';
  import useTablePagination from '../../../../utils/hooks/use-table-pagination';

  const { pagination, updatePagination } = useTablePagination('dataSource');
  const tableLoading = ref(false);
  const tableData = ref([
    {
      name: '张三',
      memo: '123',
      config: 1,
      updator: '张三',
      updatedAt: '2022-01-01',
    },
    {
      name: '李四',
      memo: '123',
      config: 1,
      updator: '张三',
      updatedAt: '2022-01-01',
    },
    {
      name: '王五',
      memo: '123',
      config: 0,
      updator: '张三',
      updatedAt: '2022-01-01',
    },
  ]);

  const handleEditData = (tableItem: any) => {
    console.log(tableItem);
  };
  const handleDeleteDataSource = (tableItem: any) => {
    console.log(tableItem);
  };

  const handlePageLimitChange = (val: number) => {
    updatePagination('limit', val);
  };

  const handlePageCurrentChange = (val: number) => {
    pagination.value.current = val;
  };

  const handleImportTable = (tableItem: any) => {
    console.log(tableItem);
  };

  const handleExportTable = (tableItem: any) => {
    console.log(tableItem);
  };

  const handleDeleteTable = (tableItem: any) => {
    console.log(tableItem);
  };
</script>

<style scoped lang="scss">
  .action-btns {
    display: flex;
    align-items: center;
    height: 100%;
    display: flex;
    align-items: center;
    height: 100%;
    .more-actions {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-left: 8px;
      width: 16px;
      height: 16px;
      border-radius: 50%;
      cursor: pointer;
      &:hover {
        background: #dcdee5;
        color: #3a84ff;
      }
    }
    .ellipsis-icon {
      transform: rotate(90deg);
    }
    .bk-button {
      margin-right: 8px;
    }
  }
</style>

<style lang="scss">
  .trusteeship-table-actions-popover.bk-popover.bk-pop2-content {
    padding: 4px 0;
    border: 1px solid #dcdee5;
    box-shadow: 0 2px 6px 0 #0000001a;
    .config-actions {
      .action-item {
        padding: 0 12px;
        min-width: 58px;
        height: 32px;
        line-height: 32px;
        color: #63656e;
        font-size: 12px;
        cursor: pointer;
        &:hover {
          background: #f5f7fa;
        }

        &.disabled {
          color: #c4c6cc;
          cursor: not-allowed;
        }
      }
    }
  }
</style>
