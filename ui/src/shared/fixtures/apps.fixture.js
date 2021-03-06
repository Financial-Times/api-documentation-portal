export default {
  "services": [
    {
      "name": "draft-content-api",
      "title": "Draft Content API",
      "description": "API for draft content",
      "tags": [
        "Draft Content",
        "Info",
        "Health"
      ],
      "paths": [
        "/drafts/content"
      ],
      "api": "https://pre-prod-eu-pac.ft.com/__api-documentation-portal/services/draft-content-api/__api"
    },
    {
      "name": "annotations-publisher",
      "title": "Annotations Publisher",
      "description": "Publishes annotations to UPP from PAC",
      "tags": [
        "Public API",
        "Health",
        "Info"
      ],
      "paths": [
        "/drafts/annotations"
      ],
      "api": "https://pre-prod-eu-pac.ft.com/__api-documentation-portal/services/annotations-publisher/__api"
    }
  ]
}
