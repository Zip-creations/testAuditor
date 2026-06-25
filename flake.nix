{
  description = "testAuditor Go tool";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "testAuditor";
        version = "0.1.0";

        src = ./.;

        vendorHash = null;

        subPackages = [ "./src" ];
      };

      apps.${system}.default = {
        type = "app";
        program = "${self.packages.${system}.default}/bin/testAuditor";
      };

      devShells.${system}.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go
        ];
      };
    };
}
