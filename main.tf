terraform {
  required_providers {
    filr = {
      source = "richardstrnad/filr"
    }
  }
}

provider "filr" {
  folder = "data"
}

resource "filr_file_state" "file1" {
  content = "File content"
}

resource "filr_file_state" "file2" {
  content = "Other file content"
}
