{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = pkgs.mkShell {
          packages = with pkgs; [
            gcc
            go
            mecab
            mozcdic-ut-neologd
          ];
          shellHook = ''
            export CGO_ENABLED=1
            export CGO_LDFLAGS="-L${pkgs.mecab}/lib -lmecab"
            export CGO_CFLAGS="-I${pkgs.mecab}/include"
          '';
        };
      }
    );
}
