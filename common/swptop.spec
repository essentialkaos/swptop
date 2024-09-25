################################################################################

%define debug_package  %{nil}

################################################################################

Summary:         Utility for viewing swap consumption of processes
Name:            swptop
Version:         1.1.0
Release:         0%{?dist}
Group:           Applications/System
License:         Apache License, Version 2.0
URL:             https://kaos.sh/swptop

Source0:         https://source.kaos.st/%{name}/%{name}-%{version}.tar.bz2

BuildRoot:       %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

BuildRequires:   golang >= 1.22

Provides:        %{name} = %{version}-%{release}

################################################################################

%description
Utility for viewing swap consumption of processes.

################################################################################

%prep

%setup -q
if [[ ! -d "%{name}/vendor" ]] ; then
  echo -e "----\nThis package requires vendored dependencies\n----"
  exit 1
elif [[ -f "%{name}/%{name}" ]] ; then
  echo -e "----\nSources must not contain precompiled binaries\n----"
  exit 1
fi

%build
pushd %{name}
  %{__make} %{?_smp_mflags} all
  cp LICENSE ..
popd

%install
rm -rf %{buildroot}

install -dm 755 %{buildroot}%{_bindir}
install -dm 755 %{buildroot}%{_mandir}/man1

install -pm 755 %{name}/%{name} %{buildroot}%{_bindir}/

./%{name}/%{name} --generate-man > %{buildroot}%{_mandir}/man1/%{name}.1

%clean
rm -rf %{buildroot}

%post
if [[ -d %{_sysconfdir}/bash_completion.d ]] ; then
  %{name} --completion=bash 1> %{_sysconfdir}/bash_completion.d/%{name} 2>/dev/null
fi

if [[ -d %{_datarootdir}/fish/vendor_completions.d ]] ; then
  %{name} --completion=fish 1> %{_datarootdir}/fish/vendor_completions.d/%{name}.fish 2>/dev/null
fi

if [[ -d %{_datadir}/zsh/site-functions ]] ; then
  %{name} --completion=zsh 1> %{_datadir}/zsh/site-functions/_%{name} 2>/dev/null
fi

%postun
if [[ $1 == 0 ]] ; then
  if [[ -f %{_sysconfdir}/bash_completion.d/%{name} ]] ; then
    rm -f %{_sysconfdir}/bash_completion.d/%{name} &>/dev/null || :
  fi

  if [[ -f %{_datarootdir}/fish/vendor_completions.d/%{name}.fish ]] ; then
    rm -f %{_datarootdir}/fish/vendor_completions.d/%{name}.fish &>/dev/null || :
  fi

  if [[ -f %{_datadir}/zsh/site-functions/_%{name} ]] ; then
    rm -f %{_datadir}/zsh/site-functions/_%{name} &>/dev/null || :
  fi
fi

################################################################################

%files
%defattr(-,root,root,-)
%doc LICENSE
%{_mandir}/man1/%{name}.1.*
%{_bindir}/%{name}

################################################################################

%changelog
* Wed Sep 25 2024 Anton Novojilov <andy@essentialkaos.com> - 1.1.0-0
- Add more info to verbose version info
- Dependencies update

* Mon Jun 24 2024 Anton Novojilov <andy@essentialkaos.com> - 1.0.1-0
- Code refactoring
- Dependencies update

* Thu Mar 28 2024 Anton Novojilov <andy@essentialkaos.com> - 1.0.0-0
- Improved support information gathering
- Code refactoring
- Dependencies update

* Sun Feb 26 2023 Anton Novojilov <andy@essentialkaos.com> - 0.6.4-0
- Added verbose version output
- Dependencies update
- Code refactoring

* Thu Dec 01 2022 Anton Novojilov <andy@essentialkaos.com> - 0.6.3-1
- Fixed build using sources from source.kaos.st

* Tue Mar 29 2022 Anton Novojilov <andy@essentialkaos.com> - 0.6.3-0
- Removed pkg.re usage
- Added module info
- Added Dependabot configuration

* Fri Dec 04 2020 Anton Novojilov <andy@essentialkaos.com> - 0.6.2-0
- ek package updated to the latest stable version

* Tue Oct 22 2019 Anton Novojilov <andy@essentialkaos.com> - 0.6.1-0
- ek package updated to the latest stable version

* Sat Jun 15 2019 Anton Novojilov <andy@essentialkaos.com> - 0.6.0-0
- ek package updated to the latest stable version
- Added completion generation for bash, zsh and fish

* Fri Sep 14 2018 Anton Novojilov <andy@essentialkaos.com> - 0.5.1-0
- Minor UI bugfixes

* Wed May 16 2018 Anton Novojilov <andy@essentialkaos.com> - 0.5.0-0
- Improved process of swap info collecting
- Fixed bug with output info if swap disabled on system

* Wed Jan 31 2018 Anton Novojilov <andy@essentialkaos.com> - 0.4.0-0
- Improved swap statistics output

* Fri Jan 12 2018 Anton Novojilov <andy@essentialkaos.com> - 0.3.1-0
- Added usage examples

* Tue Dec 19 2017 Anton Novojilov <andy@essentialkaos.com> - 0.3.0-0
- Added output filtering feature
- Output overall swap usage info
- ek package updated to latest stable release

* Fri May 26 2017 Anton Novojilov <andy@essentialkaos.com> - 0.2.0-0
- ek package updated to v9

* Fri Apr 21 2017 Anton Novojilov <andy@essentialkaos.com> - 0.1.1-0
- Added build tag

* Thu Apr 20 2017 Anton Novojilov <andy@essentialkaos.com> - 0.1.0-0
- Initial build
