{nixpkgs ? import <nixpkgs> {}}:
with nixpkgs;

mkShell {
  buildInputs = [
    (callPackage ./default.nix {})
  ];
}
