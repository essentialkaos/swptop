###############################################################################

# rpmbuilder:relative-pack true

###############################################################################

%define  debug_package %{nil}

###############################################################################

Summary:         Utility for viewing swap consumption of processes
Name:            swptop
Version:         0.2.0
Release:         0%{?dist}
Group:           Applications/System
License:         EKOL
URL:             https://github.com/essentialkaos/swptop

Source0:         https://source.kaos.io/%{name}/%{name}-%{version}.tar.bz2

BuildRoot:       %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

BuildRequires:   golang >= 1.8

Provides:        %{name} = %{version}-%{release}

###############################################################################

%description
Utility for viewing swap consumption of processes.

###############################################################################

%prep
%setup -q

%build
export GOPATH=$(pwd) 
go build src/github.com/essentialkaos/%{name}/%{name}.go

%install
rm -rf %{buildroot}

install -dm 755 %{buildroot}%{_bindir}
install -pm 755 %{name} %{buildroot}%{_bindir}/

%clean
rm -rf %{buildroot}

###############################################################################

%files
%defattr(-,root,root,-)
%doc LICENSE.EN LICENSE.RU
%{_bindir}/%{name}

###############################################################################

%changelog
* Fri May 26 2017 Anton Novojilov <andy@essentialkaos.com> - 0.2.0-0
- ek package updated to v9

* Fri Apr 21 2017 Anton Novojilov <andy@essentialkaos.com> - 0.1.1-0
- Added build tag

* Thu Apr 20 2017 Anton Novojilov <andy@essentialkaos.com> - 0.1.0-0
- Initial build
