{ pkgs ? import <nixpkgs> {} }:

with pkgs;

let
  inherit (pkgs) go;
  ai_ml = import ../ai_ml { inherit go; };
  api_gateway = import ../api_gateway { inherit go; };
  # Include other modules similarly
in
pkgs.mkShell {
  buildInputs = [ ai_ml api_gateway ];
  # Add other microservices to buildInputs
}
