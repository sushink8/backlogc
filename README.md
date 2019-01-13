# 事前準備
backlogのAPIキーを取得します  
1. backlogにログイン
1. 個人設定→API
1. APIを発行

APIを任意のファイルに保存します  

# 使い方
## プロジェクトの一覧
`./backlog -F (APIキーファイル名) -A list_project`

## 任意のプロジェクトの課題種別一覧
`./backlog -F (APIキーファイル名) -A list_issueTypes -C (プロジェクトID)`
プロジェクトID: 数値です

## 任意のプロジェクトのカテゴリー一覧
`./backlog -F (APIキーファイル名) -A list_categories -C (プロジェクトID)`
プロジェクトID: 数値です

## 課題の追加
`./backlog -F (APIキーファイル名) -A add_issues --full -C "(パラメーター)"`
### パラメーター
以下の内容をコンマ区切りで追加
- projectId(必須)
- summary(必須)
- parentIssueId
- description
- startDate
- dueDate
- estimatedHours
- actualHours
- issueTypeId(必須)
- categoryId
- versionId
- milestoneId
- priorityId(必須) (2,3,4)
- assigneeId
- notifiedUserId
- attachmentId
