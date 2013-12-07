/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 25.11.2013 21:51
 * description: Mostly 1:1 bindings to the functions defined in rpmtag.h
 */


package rpm


// #cgo LDFLAGS: -lrpm
// #include <rpm/rpmlib.h>
import "C"


type RpmTag int

const (
	RPMTAG_NOT_FOUND         = RpmTag(C.RPMTAG_NOT_FOUND)         // unknown tag
	RPMTAG_HEADERIMAGE       = RpmTag(C.RPMTAG_HEADERIMAGE)       // current image
	RPMTAG_HEADERSIGNATURES  = RpmTag(C.RPMTAG_HEADERSIGNATURES)  // signatures
	RPMTAG_HEADERIMMUTABLE   = RpmTag(C.RPMTAG_HEADERIMMUTABLE)   // original image
	RPMTAG_HEADERREGIONS     = RpmTag(C.RPMTAG_HEADERREGIONS)     // regions
	// s[] I18N string locales
	RPMTAG_HEADERI18NTABLE   = RpmTag(C.RPMTAG_HEADERI18NTABLE) 
	RPMTAG_SIG_BASE          = RpmTag(C.RPMTAG_SIG_BASE)
	RPMTAG_SIGSIZE           = RpmTag(C.RPMTAG_SIGSIZE)           // i
	RPMTAG_SIGPGP            = RpmTag(C.RPMTAG_SIGPGP)            // x
	RPMTAG_SIGMD5            = RpmTag(C.RPMTAG_SIGMD5)            // x
	RPMTAG_PKGID             = RpmTag(C.RPMTAG_PKGID)             // x
	RPMTAG_SIGGPG            = RpmTag(C.RPMTAG_SIGGPG)            // x
	RPMTAG_PUBKEYS           = RpmTag(C.RPMTAG_PUBKEYS)           // s[]
	RPMTAG_DSAHEADER         = RpmTag(C.RPMTAG_DSAHEADER)         // x
	RPMTAG_RSAHEADER         = RpmTag(C.RPMTAG_RSAHEADER)         // x
	RPMTAG_SHA1HEADER        = RpmTag(C.RPMTAG_SHA1HEADER)        // s
	RPMTAG_HDRID             = RpmTag(C.RPMTAG_HDRID)             // s
	RPMTAG_LONGSIGSIZE       = RpmTag(C.RPMTAG_LONGSIGSIZE)       // l
	RPMTAG_LONGARCHIVESIZE   = RpmTag(C.RPMTAG_LONGARCHIVESIZE)   // l
	RPMTAG_NAME              = RpmTag(C.RPMTAG_NAME)              // s
	RPMTAG_N                 = RpmTag(C.RPMTAG_N)                 // s
	RPMTAG_VERSION           = RpmTag(C.RPMTAG_VERSION)           // s
	RPMTAG_V                 = RpmTag(C.RPMTAG_V)                 // s
	RPMTAG_RELEASE           = RpmTag(C.RPMTAG_RELEASE)           // s
	RPMTAG_R                 = RpmTag(C.RPMTAG_R)                 // s
	RPMTAG_EPOCH             = RpmTag(C.RPMTAG_EPOCH)             // i
	RPMTAG_E                 = RpmTag(C.RPMTAG_E)                 // i
	RPMTAG_SUMMARY           = RpmTag(C.RPMTAG_SUMMARY)           // s{}
	RPMTAG_DESCRIPTION       = RpmTag(C.RPMTAG_DESCRIPTION)       // s{}
	RPMTAG_BUILDTIME         = RpmTag(C.RPMTAG_BUILDTIME)         // i
	RPMTAG_BUILDHOST         = RpmTag(C.RPMTAG_BUILDHOST)         // s
	RPMTAG_INSTALLTIME       = RpmTag(C.RPMTAG_INSTALLTIME)       // i
	RPMTAG_SIZE              = RpmTag(C.RPMTAG_SIZE)              // i
	RPMTAG_DISTRIBUTION      = RpmTag(C.RPMTAG_DISTRIBUTION)      // i
	RPMTAG_VENDOR            = RpmTag(C.RPMTAG_VENDOR)            // s
	RPMTAG_GIF               = RpmTag(C.RPMTAG_GIF)               // x
	RPMTAG_XPM               = RpmTag(C.RPMTAG_XPM)               // x
	RPMTAG_LICENSE           = RpmTag(C.RPMTAG_LICENSE)           // s
	RPMTAG_PACKAGER          = RpmTag(C.RPMTAG_PACKAGER)          // s
	RPMTAG_GROUP             = RpmTag(C.RPMTAG_GROUP)             // s{}
	RPMTAG_CHANGELOG         = RpmTag(C.RPMTAG_CHANGELOG)         // s[] internal
	RPMTAG_SOURCE            = RpmTag(C.RPMTAG_SOURCE)            // s[]
	RPMTAG_PATCH             = RpmTag(C.RPMTAG_PATCH)             // s[]
	RPMTAG_URL               = RpmTag(C.RPMTAG_URL)               // s
	RPMTAG_OS                = RpmTag(C.RPMTAG_OS)                // s legacy used int
	RPMTAG_ARCH              = RpmTag(C.RPMTAG_ARCH)              // s legacy used int
	RPMTAG_PREIN             = RpmTag(C.RPMTAG_PREIN)             // s
	RPMTAG_POSTIN            = RpmTag(C.RPMTAG_POSTIN)            // s
	RPMTAG_PREUN             = RpmTag(C.RPMTAG_PREUN)             // s
	RPMTAG_POSTUN            = RpmTag(C.RPMTAG_POSTUN)            // s
	RPMTAG_FILESIZES         = RpmTag(C.RPMTAG_FILESIZES)         // i[]
	RPMTAG_FILESTATES        = RpmTag(C.RPMTAG_FILESTATES)        // c[]
	RPMTAG_FILEMODES         = RpmTag(C.RPMTAG_FILEMODES)         // h[]
	RPMTAG_FILERDEVS         = RpmTag(C.RPMTAG_FILERDEVS)         // h[]
	RPMTAG_FILEMTIMES        = RpmTag(C.RPMTAG_FILEMTIMES)        // i[]
	RPMTAG_FILEDIGESTS       = RpmTag(C.RPMTAG_FILEDIGESTS)       // s[]
	RPMTAG_FILEMD5S          = RpmTag(C.RPMTAG_FILEMD5S)          // s[]
	RPMTAG_FILELINKTOS       = RpmTag(C.RPMTAG_FILELINKTOS)       // s[]
	RPMTAG_FILEFLAGS         = RpmTag(C.RPMTAG_FILEFLAGS)         // i[]
	RPMTAG_FILEUSERNAME      = RpmTag(C.RPMTAG_FILEUSERNAME)      // s[]
	RPMTAG_FILEGROUPNAME     = RpmTag(C.RPMTAG_FILEGROUPNAME)     // s[]
	RPMTAG_ICON              = RpmTag(C.RPMTAG_ICON)              // x
	RPMTAG_SOURCERPM         = RpmTag(C.RPMTAG_SOURCERPM)         // s
	RPMTAG_FILEVERIFYFLAGS   = RpmTag(C.RPMTAG_FILEVERIFYFLAGS)   // i[]
	RPMTAG_ARCHIVESIZE       = RpmTag(C.RPMTAG_ARCHIVESIZE)       // i
	RPMTAG_PROVIDENAME       = RpmTag(C.RPMTAG_PROVIDENAME)       // s[]
	RPMTAG_PROVIDES          = RpmTag(C.RPMTAG_PROVIDES)          // s[]
	RPMTAG_P                 = RpmTag(C.RPMTAG_P)                 // s[]
	RPMTAG_REQUIREFLAGS      = RpmTag(C.RPMTAG_REQUIREFLAGS)      // i[]
	RPMTAG_REQUIRENAME       = RpmTag(C.RPMTAG_REQUIRENAME)       // s[]
	RPMTAG_REQUIRES          = RpmTag(C.RPMTAG_REQUIRES)          // s[]
	RPMTAG_REQUIREVERSION    = RpmTag(C.RPMTAG_REQUIREVERSION)    // s[]
	RPMTAG_NOSOURCE          = RpmTag(C.RPMTAG_NOSOURCE)          // i
	RPMTAG_NOPATCH           = RpmTag(C.RPMTAG_NOPATCH)           // i
	RPMTAG_CONFLICTFLAGS     = RpmTag(C.RPMTAG_CONFLICTFLAGS)     // i[]
	RPMTAG_CONFLICTNAME      = RpmTag(C.RPMTAG_CONFLICTNAME)      // s[]
	RPMTAG_CONFLICTS         = RpmTag(C.RPMTAG_CONFLICTS)         // s[]
	RPMTAG_C                 = RpmTag(C.RPMTAG_C)                 // s[]
	RPMTAG_CONFLICTVERSION   = RpmTag(C.RPMTAG_CONFLICTVERSION)   // s[]
	RPMTAG_EXCLUDEARCH       = RpmTag(C.RPMTAG_EXCLUDEARCH)       // s[]
	RPMTAG_EXCLUDEOS         = RpmTag(C.RPMTAG_EXCLUDEOS)         // s[]
	RPMTAG_EXCLUSIVEARCH     = RpmTag(C.RPMTAG_EXCLUSIVEARCH)     // s[]
	RPMTAG_EXCLUSIVEOS       = RpmTag(C.RPMTAG_EXCLUSIVEOS)       // s[]
	RPMTAG_RPMVERSION        = RpmTag(C.RPMTAG_RPMVERSION)        // s
	RPMTAG_TRIGGERSCRIPTS    = RpmTag(C.RPMTAG_TRIGGERSCRIPTS)    // s[]
	RPMTAG_TRIGGERNAME       = RpmTag(C.RPMTAG_TRIGGERNAME)       // s[]
	RPMTAG_TRIGGERVERSION    = RpmTag(C.RPMTAG_TRIGGERVERSION)    // s[]
	RPMTAG_TRIGGERFLAGS      = RpmTag(C.RPMTAG_TRIGGERFLAGS)      // i[]
	RPMTAG_TRIGGERINDEX      = RpmTag(C.RPMTAG_TRIGGERINDEX)      // i[]
	RPMTAG_VERIFYSCRIPT      = RpmTag(C.RPMTAG_VERIFYSCRIPT)      // s
	RPMTAG_CHANGELOGTIME     = RpmTag(C.RPMTAG_CHANGELOGTIME)     // i[]
	RPMTAG_CHANGELOGNAME     = RpmTag(C.RPMTAG_CHANGELOGNAME)     // s[]
	RPMTAG_CHANGELOGTEXT     = RpmTag(C.RPMTAG_CHANGELOGTEXT)     // s[]
	RPMTAG_PREINPROG         = RpmTag(C.RPMTAG_PREINPROG)         // s[]
	RPMTAG_POSTINPROG        = RpmTag(C.RPMTAG_POSTINPROG)        // s[]
	RPMTAG_PREUNPROG         = RpmTag(C.RPMTAG_PREUNPROG)         // s[]
	RPMTAG_POSTUNPROG        = RpmTag(C.RPMTAG_POSTUNPROG)        // s[]
	RPMTAG_BUILDARCHS        = RpmTag(C.RPMTAG_BUILDARCHS)        // s[]
	RPMTAG_OBSOLETENAME      = RpmTag(C.RPMTAG_OBSOLETENAME)      // s[]
	RPMTAG_OBSOLETES         = RpmTag(C.RPMTAG_OBSOLETES)         // s[]
	RPMTAG_O                 = RpmTag(C.RPMTAG_O)                 // s[]
	RPMTAG_VERIFYSCRIPTPROG  = RpmTag(C.RPMTAG_VERIFYSCRIPTPROG)  // s[]
	RPMTAG_TRIGGERSCRIPTPROG = RpmTag(C.RPMTAG_TRIGGERSCRIPTPROG) // s[]
	RPMTAG_COOKIE            = RpmTag(C.RPMTAG_COOKIE)            // s
	RPMTAG_FILEDEVICES       = RpmTag(C.RPMTAG_FILEDEVICES)       // i[]
	RPMTAG_FILEINODES        = RpmTag(C.RPMTAG_FILEINODES)        // i[]
	RPMTAG_FILELANGS         = RpmTag(C.RPMTAG_FILELANGS)         // s[]
	RPMTAG_PREFIXES          = RpmTag(C.RPMTAG_PREFIXES)          // s[]
	RPMTAG_INSTPREFIXES      = RpmTag(C.RPMTAG_INSTPREFIXES)      // s[]
	RPMTAG_SOURCEPACKAGE     = RpmTag(C.RPMTAG_SOURCEPACKAGE)     // i
	RPMTAG_PROVIDEFLAGS      = RpmTag(C.RPMTAG_PROVIDEFLAGS)      // i[]
	RPMTAG_PROVIDEVERSION    = RpmTag(C.RPMTAG_PROVIDEVERSION)    // s[]
	RPMTAG_OBSOLETEFLAGS     = RpmTag(C.RPMTAG_OBSOLETEFLAGS)     // i[]
	RPMTAG_OBSOLETEVERSION   = RpmTag(C.RPMTAG_OBSOLETEVERSION)   // s[]
	RPMTAG_DIRINDEXES        = RpmTag(C.RPMTAG_DIRINDEXES)        // i[]
	RPMTAG_BASENAMES         = RpmTag(C.RPMTAG_BASENAMES)         // s[]
	RPMTAG_DIRNAMES          = RpmTag(C.RPMTAG_DIRNAMES)          // s[]
	RPMTAG_ORIGDIRINDEXES    = RpmTag(C.RPMTAG_ORIGDIRINDEXES)    // i[] relocation
	RPMTAG_ORIGBASENAMES     = RpmTag(C.RPMTAG_ORIGBASENAMES)     // s[] relocation
	RPMTAG_ORIGDIRNAMES      = RpmTag(C.RPMTAG_ORIGDIRNAMES)      // s[] relocation
	RPMTAG_OPTFLAGS          = RpmTag(C.RPMTAG_OPTFLAGS)          // s
	RPMTAG_DISTURL           = RpmTag(C.RPMTAG_DISTURL)           // s
	RPMTAG_PAYLOADFORMAT     = RpmTag(C.RPMTAG_PAYLOADFORMAT)     // s
	RPMTAG_PAYLOADCOMPRESSOR = RpmTag(C.RPMTAG_PAYLOADCOMPRESSOR) // s
	RPMTAG_PAYLOADFLAGS      = RpmTag(C.RPMTAG_PAYLOADFLAGS)      // s
	// i transaction color when installed
	RPMTAG_INSTALLCOLOR      = RpmTag(C.RPMTAG_INSTALLCOLOR)
	RPMTAG_INSTALLTID        = RpmTag(C.RPMTAG_INSTALLTID)        // i
	RPMTAG_REMOVETID         = RpmTag(C.RPMTAG_REMOVETID)         // i
	RPMTAG_PLATFORM          = RpmTag(C.RPMTAG_PLATFORM)          // s
	RPMTAG_FILECOLORS        = RpmTag(C.RPMTAG_FILECOLORS)        // i[]
	RPMTAG_FILECLASS         = RpmTag(C.RPMTAG_FILECLASS)         // i[]
	RPMTAG_CLASSDICT         = RpmTag(C.RPMTAG_CLASSDICT)         // s[]
	RPMTAG_FILEDEPENDSX      = RpmTag(C.RPMTAG_FILEDEPENDSX)      // i[]
	RPMTAG_FILEDEPENDSN      = RpmTag(C.RPMTAG_FILEDEPENDSN)      // i[]
	RPMTAG_DEPENDSDICT       = RpmTag(C.RPMTAG_DEPENDSDICT)       // i[]
	RPMTAG_SOURCEPKGID       = RpmTag(C.RPMTAG_SOURCEPKGID)       // x
	RPMTAG_FSCONTEXTS        = RpmTag(C.RPMTAG_FSCONTEXTS)        // s[] extension
	RPMTAG_RECONTEXTS        = RpmTag(C.RPMTAG_RECONTEXTS)        // s[] extension
	// s[] selinux *.te policy file
	RPMTAG_POLICIES          = RpmTag(C.RPMTAG_POLICIES)
	RPMTAG_PRETRANS          = RpmTag(C.RPMTAG_PRETRANS)          // s
	RPMTAG_POSTTRANS         = RpmTag(C.RPMTAG_POSTTRANS)         // s
	RPMTAG_PRETRANSPROG      = RpmTag(C.RPMTAG_PRETRANSPROG)      // s[]
	RPMTAG_POSTTRANSPROG     = RpmTag(C.RPMTAG_POSTTRANSPROG)     // s[]
	RPMTAG_DISTTAG           = RpmTag(C.RPMTAG_DISTTAG)           // s
	RPMTAG_DBINSTANCE        = RpmTag(C.RPMTAG_DBINSTANCE)        // i extension
	RPMTAG_NVRA              = RpmTag(C.RPMTAG_NVRA)              // s extension
	RPMTAG_FILENAMES         = RpmTag(C.RPMTAG_FILENAMES)         // s[] extension
	RPMTAG_FILEPROVIDE       = RpmTag(C.RPMTAG_FILEPROVIDE)       // s[] extension
	RPMTAG_FILEREQUIRE       = RpmTag(C.RPMTAG_FILEREQUIRE)       // s[] extension
	RPMTAG_TRIGGERCONDS      = RpmTag(C.RPMTAG_TRIGGERCONDS)      // s[] extension
	RPMTAG_TRIGGERTYPE       = RpmTag(C.RPMTAG_TRIGGERTYPE)       // s[] extension
	RPMTAG_ORIGFILENAMES     = RpmTag(C.RPMTAG_ORIGFILENAMES)     // s[] extension
	RPMTAG_LONGFILESIZES     = RpmTag(C.RPMTAG_LONGFILESIZES)     // l[]
	RPMTAG_LONGSIZE          = RpmTag(C.RPMTAG_LONGSIZE)          // l
	RPMTAG_FILECAPS          = RpmTag(C.RPMTAG_FILECAPS)          // s[]
	// i file digest algorithm
	RPMTAG_FILEDIGESTALGO    = RpmTag(C.RPMTAG_FILEDIGESTALGO)
	RPMTAG_BUGURL            = RpmTag(C.RPMTAG_BUGURL)            // s
	RPMTAG_EVR               = RpmTag(C.RPMTAG_EVR)               // s extension
	RPMTAG_NVR               = RpmTag(C.RPMTAG_NVR)               // s extension
	RPMTAG_NEVR              = RpmTag(C.RPMTAG_NEVR)              // s extension
	RPMTAG_NEVRA             = RpmTag(C.RPMTAG_NEVRA)             // s extension
	RPMTAG_HEADERCOLOR       = RpmTag(C.RPMTAG_HEADERCOLOR)       // i extension
	RPMTAG_VERBOSE           = RpmTag(C.RPMTAG_VERBOSE)           // i extension
	RPMTAG_EPOCHNUM          = RpmTag(C.RPMTAG_EPOCHNUM)          // i extension 
)

