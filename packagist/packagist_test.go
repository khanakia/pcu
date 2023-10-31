package packagist

import (
	"testing"
)

func Test_LatestPackageVersion(t *testing.T) {
	repo, _ := OpenRepoFile("./test/lighthouse.json")
	pkgs := GetPackagesFromList(repo.PackagesMap)
	pkg := GetLatestPackage(pkgs)
	if pkg.VersionNormalized != "6.22.0.0" {
		t.Errorf("returned = %s; want 6.22.0.0", pkg.VersionNormalized)
	}
}

func Test_LatestPackageVersionWithBeta(t *testing.T) {
	repo, _ := OpenRepoFile("./test/dbal.json")
	pkgs := GetPackagesFromList(repo.PackagesMap)
	pkg := GetLatestPackage(pkgs)
	if pkg.VersionNormalized != "3.7.1.0" {
		t.Errorf("returned = %s; want 3.7.1.0", pkg.VersionNormalized)
	}
}

func Test_DevMode(t *testing.T) {
	ok := IsDevelopmentPackage("4.0.0.0-RC1")
	if !ok {
		t.Errorf("returned false for 4.0.0.0-RC1 wants true")
	}
}
