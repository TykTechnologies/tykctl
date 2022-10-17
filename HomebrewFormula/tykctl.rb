# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
require_relative "lib/custom_download_strategy"
class Tykctl < Formula
  desc ""
  homepage ""
  version "0.0.12"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.12/tykctl_0.0.12_Darwin_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "ff39da5833d0761964fe37446f05f438775d5ba3f4c11e385d98ca0650fd9fe1"

      def install
        bin.install "tykctl"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.12/tykctl_0.0.12_Darwin_x86_64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "e0ecf1112ab88446344b13acdf3a8ce47e2a033bb21f3453bdffa38362990e9e"

      def install
        bin.install "tykctl"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.12/tykctl_0.0.12_Linux_x86_64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "b18da2c164a456a32537eb95915fa52d5e9692e3b3b12da2ace0ed807b680b50"

      def install
        bin.install "tykctl"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.12/tykctl_0.0.12_Linux_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "386f520dbb2385683c93543797ea2c048f4ce0de9d30bb31ba62e43ba6b3c66b"

      def install
        bin.install "tykctl"
      end
    end
  end
end
