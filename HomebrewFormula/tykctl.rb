# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
require_relative "lib/custom_download_strategy"
class Tykctl < Formula
  desc ""
  homepage ""
  version "0.0.25"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.25/tykctl_0.0.25_Darwin_x86_64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "e81981d06a70d95a98a6ee792cc91e9bc1baed82cd2fc5d6284a5933269ef8f2"

      def install
        bin.install "tykctl"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.25/tykctl_0.0.25_Darwin_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "bc83bc5f7d13654dd07775c1bbfd7bad3ee36bf79efc2f51aae811f1356e011a"

      def install
        bin.install "tykctl"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.25/tykctl_0.0.25_Linux_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "071d3b671bf10a3b782ed9fe8a94ed9a4c79a1beff153eaabb884f2d94c464f1"

      def install
        bin.install "tykctl"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.25/tykctl_0.0.25_Linux_x86_64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "2203e7f6e5f63864a11a5f17d8b0331451dd8b05dae978f501f2b205073a5db0"

      def install
        bin.install "tykctl"
      end
    end
  end
end
