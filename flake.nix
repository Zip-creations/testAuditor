{
  description = "testAuditor Go tool";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs = inputs@{ self, nixpkgs, flake-parts, ... }:
      flake-parts.lib.mkFlake { inherit inputs; } {
        systems = [
          "x86_64-linux"
          "x86_64-darwin"
          "aarch64-darwin"
          "aarch64-linux"
        ];
        
        perSystem =
          { system, self', ... }:
          let
            pkgs = import nixpkgs { inherit system; };
          in
            {
              packages.default = pkgs.buildGoModule {
                pname = "testAuditor";
                version = "0.1.0";
                
                src = ./src;
                
                vendorHash = null;
                
                subPackages = [ "cmd/testAuditor" ];
              };

              apps.default = {
                type = "app";
                program = "${self.packages.${system}.default}/bin/testAuditor";
              };
              
              devShells.default = pkgs.mkShell {
                buildInputs = [
                  pkgs.go
                  pkgs.gopls
                ];
              };
            };
      };
}
