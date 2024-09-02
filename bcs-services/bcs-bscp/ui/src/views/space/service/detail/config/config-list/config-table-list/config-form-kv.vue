<template>
  <bk-form ref="formRef" form-type="vertical" :model="localVal" :rules="rules">
    <div class="key-and-type">
      <bk-form-item :label="t('配置项名称')" property="key" :required="true">
        <bk-input v-model="localVal.key" :disabled="props.editMode" @input="change" :placeholder="t('请输入')" />
      </bk-form-item>
      <bk-form-item :label="t('数据类型')" property="kv_type" :required="true" :description="typeDescription">
        <bk-select
          class="bk-select"
          v-model="localVal.kv_type"
          filterable
          :disabled="selectDisabled"
          :clearable="false">
          <bk-option v-for="kvType in CONFIG_KV_TYPE" :key="kvType.id" :id="kvType.id" :name="kvType.name" />
        </bk-select>
      </bk-form-item>
    </div>
    <bk-form-item :label="t('配置项描述')" property="memo">
      <bk-input v-model="localVal.memo" type="textarea" :maxlength="200" :placeholder="t('请输入')" @input="change" />
    </bk-form-item>
    <KvTaleForm v-if="localVal.kv_type === 'table'" />
    <bk-form-item v-else :label="t('配置项值')" property="value" :required="true">
      <bk-input
        v-if="localVal.kv_type === 'string' || localVal.kv_type === 'number'"
        v-model.trim="localVal!.value"
        @input="change" />
      <KvConfigContentEditor
        v-else
        ref="KvCodeEditorRef"
        :languages="localVal.kv_type"
        :content="localVal.value"
        @change="handleStringContentChange" />
    </bk-form-item>
  </bk-form>
</template>

<script lang="ts" setup>
  import { ref, onMounted, computed } from 'vue';
  import { useI18n } from 'vue-i18n';
  import { CONFIG_KV_TYPE } from '../../../../../../../constants/config';
  import KvConfigContentEditor from '../../components/kv-config-content-editor.vue';
  import { IConfigKvEditParams } from '../../../../../../../../types/config';
  import useServiceStore from '../../../../../../../store/service';
  import { storeToRefs } from 'pinia';
  import KvTaleForm from './kv-table-form.vue';

  const serviceStore = useServiceStore();
  const { appData } = storeToRefs(serviceStore);
  const { t } = useI18n();

  const props = withDefaults(
    defineProps<{
      config: IConfigKvEditParams;
      editMode?: boolean;
      bkBizId: string;
      id: number; // 服务ID或者模板空间ID
      isTpl?: boolean; // 是否未模板配置文件，非模板配置文件和模板配置文件的上传、下载接口参数有差异
    }>(),
    {
      editMode: false,
    },
  );

  const KvCodeEditorRef = ref();
  const formRef = ref();
  const localVal = ref({
    ...props.config,
  });

  const typeDescription = computed(() => {
    if (appData.value.spec.data_type !== 'any' && !props.editMode) {
      return `已限制该服务下所有配置项数据类型为${appData.value.spec.data_type}，如需其他数据类型，请调整服务属性下的数据类型`;
    }
    return '';
  });

  const selectDisabled = computed(() => appData.value.spec.data_type !== 'any' || props.editMode);

  const rules = {
    key: [
      {
        validator: (value: string) => value.length <= 128,
        message: t('最大长度128个字符'),
      },
      {
        validator: (value: string) =>
          /^[\p{Script=Han}\p{L}\p{N}]([\p{Script=Han}\p{L}\p{N}_-]*[\p{Script=Han}\p{L}\p{N}])?$/u.test(value),
        message: t('只允许包含中文、英文、数字、下划线 (_)、连字符 (-)，并且必须以中文、英文、数字开头和结尾'),
      },
    ],
    memo: [
      {
        validator: (value: string) => value.length <= 200,
        message: t('最大长度200个字符'),
      },
    ],
    value: [
      {
        validator: (value: string) => {
          if (localVal.value.kv_type === 'number') {
            return /^-?\d+(\.\d+)?$/.test(value);
          }
          return true;
        },
        message: t('配置项值不为数字'),
      },
    ],
  };

  // 新建文件任意类型默认选中string
  onMounted(() => {
    if (!props.editMode) {
      localVal.value.kv_type = appData.value.spec.data_type! === 'any' ? 'string' : appData.value.spec.data_type!;
    }
  });

  const validate = async () => {
    await formRef.value.validate();
    switch (localVal.value.kv_type) {
      case 'json':
      case 'xml':
      case 'yaml':
        return KvCodeEditorRef.value.validate();
    }
    return true;
  };

  const emits = defineEmits(['change']);

  const handleStringContentChange = (val: string) => {
    localVal.value!.value = val;
    change();
  };

  const change = () => {
    emits('change', localVal.value);
  };

  defineExpose({ validate });
</script>

<style scoped lang="scss">
  .key-and-type {
    display: flex;
    justify-content: space-between;
  }
  .bk-input,
  .bk-select {
    width: 428px;
  }
</style>
