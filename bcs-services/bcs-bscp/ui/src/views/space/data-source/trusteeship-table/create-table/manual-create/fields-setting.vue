<template>
  <Card :title="$t('字段设置')">
    <template #suffix>
      <div class="add-fields" @click="handleAddFields">
        <Plus class="add-icon" />
        <span class="text">{{ $t('添加字段') }}</span>
      </div>
    </template>
    <bk-table
      :data="tableData"
      class="fields-setting-table"
      :border="['row', 'col', 'outer']"
      row-hover="auto"
      :cell-class="getCellCls"
      :row-draggable="{ width: 20, label: '' }"
      :show-overflow-tooltip="true">
      <bk-table-column :label="$t('字段名')" :width="120">
        <template #default="{ row }">
          <bk-input v-model="row.fieldsName"></bk-input>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('显示名')" prop="key" :width="144">
        <template #default="{ row }">
          <bk-input v-model="row.showName"></bk-input>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('数据类型')" :width="126">
        <template #default="{ row }">
          <bk-select v-if="row" class="type-select" auto-focus :filterable="false" @select="row.type = $event">
            <bk-option v-for="item in dataType" :id="item.value" :key="item.value" :name="item.label" />
          </bk-select>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('默认值/枚举值')" :width="183">
        <template #default="{ row }">
          <bk-input v-if="row.type !== 'enum'" v-model="row.default"></bk-input>
          <div v-else class="enum-type">
            <bk-select class="type-select">
              <bk-option v-for="item in dataType" :id="item.value" :key="item.value" :name="item.label" />
            </bk-select>
            <bk-popover
              :width="480"
              placement="bottom-end"
              theme="light"
              trigger="click"
              ext-cls="setting-enum-popover">
              <div class="setting-icon">
                <cog-shape v-bk-tooltips="{ content: $t('设置枚举值') }" />
              </div>
              <template #content>
                <div class="setting-enum-wrap">
                  <div class="title">{{ $t('设置枚举值') }}</div>
                  <bk-radio-group v-model="settingEnumType" class="enum-radio-group">
                    <bk-radio label="single">{{ $t('单选') }}</bk-radio>
                    <bk-radio label="multi">{{ $t('多选') }}</bk-radio>
                  </bk-radio-group>
                  <div class="enum-list">
                    <div v-for="(enumItem, index) in settingEnumList" :key="enumItem.value" class="enum-item">
                      <div class="num">{{ index + 1 }}</div>
                      <bk-input v-model="enumItem.value" :placeholder="$t('实际值')"></bk-input>
                      <bk-input v-model="enumItem.text" :placeholder="$t('显示文本')"></bk-input>
                      <div class="action-btns">
                        <i class="bk-bscp-icon icon-reduce" @click="handleDelEnumItem(index)"></i>
                        <i class="bk-bscp-icon icon-add" @click="handleAddEnumItem(index)"></i>
                      </div>
                    </div>
                  </div>
                  <div class="footer">
                    <bk-button theme="primary">{{ $t('保存') }}</bk-button>
                    <bk-button>{{ $t('取消') }}</bk-button>
                  </div>
                </div>
              </template>
            </bk-popover>
          </div>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('主键')" property="primaryKey" :width="57">
        <template #default="{ row }">
          <bk-radio :checked="row.primaryKey" @change="handleChangePrimaryKey(row, $event)"></bk-radio>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('非空')" property="nonempty" :width="57">
        <bk-checkbox></bk-checkbox>
      </bk-table-column>
      <bk-table-column :label="$t('唯一')" property="only" :width="57">
        <bk-checkbox></bk-checkbox>
      </bk-table-column>
      <bk-table-column :label="$t('自增')" property="autoIncrement" :width="57">
        <bk-checkbox></bk-checkbox>
      </bk-table-column>
      <bk-table-column :label="$t('只读')" property="readonly" :width="57">
        <bk-checkbox></bk-checkbox>
      </bk-table-column>
      <bk-table-column label="" property="delete" :width="48">
        <template #default="{ row }">
          <i class="bk-bscp-icon icon-reduce delete-icon" @click="handleDelete(row.fieldsName)"></i>
        </template>
      </bk-table-column>
    </bk-table>
  </Card>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Plus, CogShape } from 'bkui-vue/lib/icon';
  import Card from '../../../component/card.vue';

  const tableData = ref<any[]>([]);
  const settingEnumType = ref('single');
  const settingEnumList = ref([{ text: '', value: '' }]);

  const dataType = [
    {
      value: 'string',
      label: 'String',
    },
    {
      value: 'number',
      label: 'Number',
    },
    {
      value: 'enum',
      label: 'ENUM',
    },
  ];

  const handleAddFields = () => {
    tableData.value.push({
      fieldsName: '',
      showName: '',
      type: '',
      required: false,
      default: '',
      primaryKey: false,
    });
  };

  const handleChangePrimaryKey = (row: any, value: boolean) => {
    console.log(row, value);
    row.primaryKey = value;
  };

  // 添加自定义单元格class
  const getCellCls = ({ property }: { property: string }) => {
    return ['primaryKey', 'nonempty', 'only', 'autoIncrement', 'readonly', 'delete'].includes(property)
      ? 'check-cell'
      : '';
  };

  const handleDelete = (fieldsName: string) => {
    tableData.value = tableData.value.filter((item: any) => item.fieldsName !== fieldsName);
  };

  const handleAddEnumItem = (index: number) => {
    settingEnumList.value.splice(index + 1, 0, { text: '', value: '' });
  };

  const handleDelEnumItem = (index: number) => {
    if (settingEnumList.value.length > 1) {
      settingEnumList.value.splice(index, 1);
    }
  };
</script>

<style scoped lang="scss">
  .add-fields {
    display: flex;
    align-items: center;
    height: 16px;
    cursor: pointer;
    .add-icon {
      border-radius: 50%;
      background-color: #3a84ff;
      color: #fff;
      margin-right: 5px;
    }
    .text {
      color: #3a84ff;
      font-size: 12px;
    }
  }
  .type-select {
    :deep(.bk-input) {
      height: 41px;
      border: none;
      .bk-input--text {
        padding-left: 16px;
      }
    }
  }
  .fields-setting-table {
    width: 100% !important;
    :deep(.bk-table-head) {
      colgroup col {
        min-width: 30px !important;
      }
    }
    :deep(.bk-table-body) {
      colgroup col {
        min-width: 30px !important;
      }
      .cell {
        padding: 0;
      }
      .check-cell {
        .cell {
          width: 100%;
          display: flex;
          align-items: center;
          justify-content: space-around;
        }
        .delete-icon {
          font-size: 16px;
          cursor: pointer;
          &:hover {
            color: #3a84ff;
          }
        }
      }
      .enum-type {
        display: flex;
        align-items: center;
        .setting-icon {
          height: 24px;
          display: flex;
          align-items: center;
          justify-content: center;
          border-left: 1px solid #dcdee5;
          width: 30px;
          height: 100%;
          font-size: 16px;
          color: #a5a8b1;
          cursor: pointer;
          &:hover {
            color: #3a84ff;
          }
        }
      }
    }
    .bk-input {
      border: none;
      height: 41px;
      &.is-focused {
        border: 1px solid #3a84ff;
      }
    }
  }

  .setting-enum-wrap {
    .title {
      font-size: 16px;
      color: #313238;
    }
    .enum-radio-group {
      margin: 12px 8px;
    }
    .enum-list {
      padding: 12px 8px;
      max-height: 300px;
      overflow: auto;
      .enum-item {
        display: flex;
        gap: 12px;
        &:not(:last-child) {
          margin-bottom: 8px;
        }
        .num {
          width: 32px;
          height: 32px;
          background: #f0f1f5;
          border-radius: 2px;
          line-height: 32px;
          text-align: center;
        }
        .bk-input {
          width: 160px;
          height: 32px;
        }
        .action-btns {
          padding: 0 8px;
          display: flex;
          align-items: center;
          gap: 8px;
          font-size: 14px;
          color: #979ba5;
          cursor: pointer;
          i:hover {
            color: #3a84ff;
          }
        }
      }
    }
    .footer {
      position: absolute;
      bottom: 0;
      left: 0;
      width: 100%;
      height: 42px;
      display: flex;
      align-items: center;
      justify-content: flex-end;
      gap: 8px;
      background: #fafbfd;
      box-shadow: 0 -1px 0 0 #dcdee5;
    }
  }
</style>

<style lang="scss">
  .bk-popover.bk-pop2-content.setting-enum-popover {
    padding-bottom: 54px;
  }
</style>
