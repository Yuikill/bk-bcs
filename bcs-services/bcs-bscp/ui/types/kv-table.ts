// 字段设置列表项
export interface IFiledsItem {
  fieldsName: string;
  showName: string;
  type: string;
  required: boolean;
  default: string;
  primaryKey: boolean;
  nonEmpty: boolean;
  only: boolean;
  autoIncrement: boolean;
  readonly: boolean;
  enumList?: IEnumItem[];
  enumType?: string;
  isShowSettingEnumPopover?: boolean;
}

// 字段设置枚举类型
export interface IEnumItem {
  text: string;
  value: string;
  hasTextError?: boolean;
  hasValueError?: boolean;
}
