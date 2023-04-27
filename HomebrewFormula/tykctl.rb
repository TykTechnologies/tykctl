# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Tykctl < Formula
  desc ""
  homepage ""
  version "0.0.29"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.29/tykctl_0.0.29_Darwin_x86_64.tar.gz"
      sha256 "525c72b02471dd8e8651ad214d93fc1c65d99df56bd5b530c735b75651a1cf89"

      def install
        bin.install "tykctl"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.29/tykctl_0.0.29_Darwin_arm64.tar.gz"
      sha256 "d5f7aa346af534e3ccac88cfa6da01a0e3507ba3153c633618bde2784e421af4"

      def install
        bin.install "tykctl"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.29/tykctl_0.0.29_Linux_arm64.tar.gz"
      sha256 "9b735632801634235a09c28c19ef51378407ac2dc31414f10f68b5425c0aa4ae"

      def install
        bin.install "tykctl"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/TykTechnologies/tykctl/releases/download/v0.0.29/tykctl_0.0.29_Linux_x86_64.tar.gz"
      sha256 "d6b1ae15eaa101878c0a3c179749ba0946723978701edc49271eac6ee08c3b33"

      def install
        bin.install "tykctl"
      end
    end
  end
end
