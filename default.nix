{ stdenv, bash }:

stdenv.mkDerivation rec {
  name = "subcommander-${version}";
  version = "1.0.0";

  src = ./.;

  propagatedBuildInputs = [
    bash
  ];

  buildPhase = "true";
  installPhase = ''
    mkdir -p $out/bin
    cp $src/subcommander $out/bin/subcommander
    chmod +x $out/bin/subcommander
  '';
}
