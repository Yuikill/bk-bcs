<template>
  <bk-form-item :label="$t('表格配置')">
    <div class="source-select">
      <div :class="['source-item', { active: sourceType === 'trusteeship' }]" @click="sourceType = 'trusteeship'">
        {{ $t('托管表格配置') }}
      </div>
      <div :class="['source-item', { active: sourceType === 'external' }]" @click="sourceType = 'external'">
        {{ $t('关联外部数据源') }}
      </div>
    </div>
    <bk-select class="table-select" :placeholder="tableSelectPlaceholder"></bk-select>
  </bk-form-item>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { useI18n } from 'vue-i18n';
  const { t } = useI18n();

  const sourceType = ref('trusteeship');

  const tableSelectPlaceholder = computed(() => {
    return sourceType.value === 'trusteeship' ? t('请选择表格') : t('请选择数据源');
  });
</script>

<style scoped lang="scss">
  .source-select {
    display: flex;
    .source-item {
      width: 214px;
      height: 26px;
      background: #ffffff;
      border: 1px solid #c4c6cc;
      border-radius: 0 2px 2px 0;
      text-align: center;
      line-height: 26px;
      font-size: 12px;
      color: #63656e;
      cursor: pointer;
      &:first-child {
        border-right: none;
        border-radius: 2px 0 0 2px;
      }
      &.active {
        background: #e1ecff;
        border: 1px solid #3a84ff;
        color: #3a84ff;
      }
    }
  }

  .table-select {
    width: 428px;
    margin-top: 12px;
  }
</style>
