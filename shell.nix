{ pkgs ? import <nixpkgs> {} }:

let
    go = pkgs.go.overrideAttrs (old: rec {
        version = "1.26.5";
        src = pkgs.fetchurl {
            url = "https://go.dev/dl/go${version}.src.tar.gz";
            hash = "sha256-SVvkvIcXasVnOS5bQRar2YRm0z17SdQedkzMaXay3EI=";
        };
    });
in
pkgs.mkShell {
    buildInputs = [ go ];
}
