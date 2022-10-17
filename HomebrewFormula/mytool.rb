# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
require_relative "lib/custom_download_strategy"
class Mytool < Formula
  desc ""
  homepage ""
  version "1.1.5"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/myorg/mytool/releases/download/v1.1.5/mytool_1.1.5_Darwin_arm64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "abc123..."

      def install
        bin.install "mytool"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/myorg/mytool/releases/download/v1.1.5/mytool_1.1.5_Darwin_x86_64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "qwerty987..."

      def install
        bin.install "mytool"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/myorg/mytool/releases/download/v1.1.5/mytool_1.1.5_Linux_arm64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "f00bar..."

      def install
        bin.install "mytool"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/myorg/mytool/releases/download/v1.1.5/mytool_1.1.5_Linux_x86_64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "xyz543..."

      def install
        bin.install "mytool"
      end
    end
  end
end