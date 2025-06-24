{
  description = "An MCP (Model Context Protocol) server to make Japan's hot summer cool";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    flake-parts.url = "github:hercules-ci/flake-parts";
    systems.url = "github:nix-systems/default";
  };

  outputs = {flake-parts, ...} @ inputs:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = import inputs.systems;

      perSystem = {
        config,
        self',
        inputs',
        pkgs,
        system,
        ...
      }: let
        pname = "summer-in-japan-is-too-hot-mcp";
        version = "0.0.0";
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            nil
            go
            gopls # Language Server for Go
          ];
        };

        packages.default = pkgs.buildGoModule {
          inherit pname version;
          src = builtins.path {
            path = ./.;
            name = "source";
          };
          vendorHash = "sha256-BTN/Tlb7q56cnWh0yjL1XeW8q12qQMSGrDmEsgGtQ2s=";
        };

        apps.stdio = {
          type = "app";
          program = "${self'.packages.default}/bin/${pname}-stdio";
        };

        apps.http = {
          type = "app";
          program = "${self'.packages.default}/bin/${pname}-http";
        };

        apps.sse = {
          type = "app";
          program = "${self'.packages.default}/bin/${pname}-sse";
        };

        apps.inprocess = {
          type = "app";
          program = "${self'.packages.default}/bin/${pname}-inprocess";
        };
      };
    };
}
